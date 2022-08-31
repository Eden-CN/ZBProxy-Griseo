package tls

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/LittleGriseo/GriseoProxy/common"
	"github.com/LittleGriseo/GriseoProxy/common/set"
	"github.com/LittleGriseo/GriseoProxy/config"
	"github.com/LittleGriseo/GriseoProxy/outbound"
	"github.com/LittleGriseo/GriseoProxy/service/access"
	"net"
)

func NewConnHandler(s *config.ConfigProxyService,
	c net.Conn,
	out outbound.Outbound) (net.Conn, error) {
	header, buf, err := SniffAndRecordTLS(c)
	if err != nil {
		if err == ErrNotTLS {
			if s.TLSSniffing.RejectNonTLS {
				buf.Reset()
				return nil, err
			}
			return dialAndWrite(s, buf, out)
		}
		return nil, err
	}
	domain := header.Domain()
	hit := false
	for _, list := range s.TLSSniffing.SNIAllowListTags {
		if hit = common.Must[*set.StringSet](access.GetTargetList(list)).Has(domain); hit {
			break
		}
	}
	if !hit {
		if s.TLSSniffing.RejectIfNonMatch {
			buf.Reset()
			return nil, errors.New("")
		}
		return dialAndWrite(s, buf, out)
	}
	defer buf.Reset()
	remote, err := out.Dial("tcp", fmt.Sprintf("%s:%v", domain, s.TargetPort))
	if err != nil {
		return nil, err
	}
	_, err = buf.WriteTo(remote)
	if err != nil {
		return nil, err
	}
	return remote, nil
}

func dialAndWrite(s *config.ConfigProxyService, buffer *bytes.Buffer, out outbound.Outbound) (net.Conn, error) {
	defer buffer.Reset()
	conn, err := out.Dial("tcp", fmt.Sprintf("%s:%v", s.TargetAddress, s.TargetPort))
	if err != nil {
		return nil, err
	}
	_, err = buffer.WriteTo(conn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
