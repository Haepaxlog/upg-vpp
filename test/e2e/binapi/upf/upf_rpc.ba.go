// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package upf

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/travelping/upg-vpp/test/e2e/binapi/memclnt"
)

// RPCService defines RPC service upf.
type RPCService interface {
	UpfAppAddDel(ctx context.Context, in *UpfAppAddDel) (*UpfAppAddDelReply, error)
	UpfAppFlowTimeoutSet(ctx context.Context, in *UpfAppFlowTimeoutSet) (*UpfAppFlowTimeoutSetReply, error)
	UpfAppIPRuleAddDel(ctx context.Context, in *UpfAppIPRuleAddDel) (*UpfAppIPRuleAddDelReply, error)
	UpfAppL7RuleAddDel(ctx context.Context, in *UpfAppL7RuleAddDel) (*UpfAppL7RuleAddDelReply, error)
	UpfApplicationL7RuleDump(ctx context.Context, in *UpfApplicationL7RuleDump) (RPCService_UpfApplicationL7RuleDumpClient, error)
	UpfApplicationsDump(ctx context.Context, in *UpfApplicationsDump) (RPCService_UpfApplicationsDumpClient, error)
	UpfNatPoolDump(ctx context.Context, in *UpfNatPoolDump) (RPCService_UpfNatPoolDumpClient, error)
	UpfNwiAddDel(ctx context.Context, in *UpfNwiAddDel) (*UpfNwiAddDelReply, error)
	UpfNwiDump(ctx context.Context, in *UpfNwiDump) (RPCService_UpfNwiDumpClient, error)
	UpfPfcpEndpointAddDel(ctx context.Context, in *UpfPfcpEndpointAddDel) (*UpfPfcpEndpointAddDelReply, error)
	UpfPfcpEndpointDump(ctx context.Context, in *UpfPfcpEndpointDump) (RPCService_UpfPfcpEndpointDumpClient, error)
	UpfPfcpFormat(ctx context.Context, in *UpfPfcpFormat) (*UpfPfcpFormatReply, error)
	UpfPfcpHeartbeatsGet(ctx context.Context, in *UpfPfcpHeartbeatsGet) (*UpfPfcpHeartbeatsGetReply, error)
	UpfPfcpHeartbeatsSet(ctx context.Context, in *UpfPfcpHeartbeatsSet) (*UpfPfcpHeartbeatsSetReply, error)
	UpfPfcpPolicerSet(ctx context.Context, in *UpfPfcpPolicerSet) (*UpfPfcpPolicerSetReply, error)
	UpfPfcpPolicerShow(ctx context.Context, in *UpfPfcpPolicerShow) (*UpfPfcpPolicerShowReply, error)
	UpfPfcpReencode(ctx context.Context, in *UpfPfcpReencode) (*UpfPfcpReencodeReply, error)
	UpfPfcpServerSet(ctx context.Context, in *UpfPfcpServerSet) (*UpfPfcpServerSetReply, error)
	UpfPfcpServerShow(ctx context.Context, in *UpfPfcpServerShow) (*UpfPfcpServerShowReply, error)
	UpfPolicyAddDel(ctx context.Context, in *UpfPolicyAddDel) (*UpfPolicyAddDelReply, error)
	UpfPolicyDump(ctx context.Context, in *UpfPolicyDump) (RPCService_UpfPolicyDumpClient, error)
	UpfUpdateApp(ctx context.Context, in *UpfUpdateApp) (*UpfUpdateAppReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) UpfAppAddDel(ctx context.Context, in *UpfAppAddDel) (*UpfAppAddDelReply, error) {
	out := new(UpfAppAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfAppFlowTimeoutSet(ctx context.Context, in *UpfAppFlowTimeoutSet) (*UpfAppFlowTimeoutSetReply, error) {
	out := new(UpfAppFlowTimeoutSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfAppIPRuleAddDel(ctx context.Context, in *UpfAppIPRuleAddDel) (*UpfAppIPRuleAddDelReply, error) {
	out := new(UpfAppIPRuleAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfAppL7RuleAddDel(ctx context.Context, in *UpfAppL7RuleAddDel) (*UpfAppL7RuleAddDelReply, error) {
	out := new(UpfAppL7RuleAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfApplicationL7RuleDump(ctx context.Context, in *UpfApplicationL7RuleDump) (RPCService_UpfApplicationL7RuleDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_UpfApplicationL7RuleDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_UpfApplicationL7RuleDumpClient interface {
	Recv() (*UpfApplicationL7RuleDetails, error)
	api.Stream
}

type serviceClient_UpfApplicationL7RuleDumpClient struct {
	api.Stream
}

func (c *serviceClient_UpfApplicationL7RuleDumpClient) Recv() (*UpfApplicationL7RuleDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *UpfApplicationL7RuleDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) UpfApplicationsDump(ctx context.Context, in *UpfApplicationsDump) (RPCService_UpfApplicationsDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_UpfApplicationsDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_UpfApplicationsDumpClient interface {
	Recv() (*UpfApplicationsDetails, error)
	api.Stream
}

type serviceClient_UpfApplicationsDumpClient struct {
	api.Stream
}

func (c *serviceClient_UpfApplicationsDumpClient) Recv() (*UpfApplicationsDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *UpfApplicationsDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) UpfNatPoolDump(ctx context.Context, in *UpfNatPoolDump) (RPCService_UpfNatPoolDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_UpfNatPoolDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_UpfNatPoolDumpClient interface {
	Recv() (*UpfNatPoolDetails, error)
	api.Stream
}

type serviceClient_UpfNatPoolDumpClient struct {
	api.Stream
}

func (c *serviceClient_UpfNatPoolDumpClient) Recv() (*UpfNatPoolDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *UpfNatPoolDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) UpfNwiAddDel(ctx context.Context, in *UpfNwiAddDel) (*UpfNwiAddDelReply, error) {
	out := new(UpfNwiAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfNwiDump(ctx context.Context, in *UpfNwiDump) (RPCService_UpfNwiDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_UpfNwiDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_UpfNwiDumpClient interface {
	Recv() (*UpfNwiDetails, error)
	api.Stream
}

type serviceClient_UpfNwiDumpClient struct {
	api.Stream
}

func (c *serviceClient_UpfNwiDumpClient) Recv() (*UpfNwiDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *UpfNwiDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) UpfPfcpEndpointAddDel(ctx context.Context, in *UpfPfcpEndpointAddDel) (*UpfPfcpEndpointAddDelReply, error) {
	out := new(UpfPfcpEndpointAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfPfcpEndpointDump(ctx context.Context, in *UpfPfcpEndpointDump) (RPCService_UpfPfcpEndpointDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_UpfPfcpEndpointDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_UpfPfcpEndpointDumpClient interface {
	Recv() (*UpfPfcpEndpointDetails, error)
	api.Stream
}

type serviceClient_UpfPfcpEndpointDumpClient struct {
	api.Stream
}

func (c *serviceClient_UpfPfcpEndpointDumpClient) Recv() (*UpfPfcpEndpointDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *UpfPfcpEndpointDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) UpfPfcpFormat(ctx context.Context, in *UpfPfcpFormat) (*UpfPfcpFormatReply, error) {
	out := new(UpfPfcpFormatReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfPfcpHeartbeatsGet(ctx context.Context, in *UpfPfcpHeartbeatsGet) (*UpfPfcpHeartbeatsGetReply, error) {
	out := new(UpfPfcpHeartbeatsGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpfPfcpHeartbeatsSet(ctx context.Context, in *UpfPfcpHeartbeatsSet) (*UpfPfcpHeartbeatsSetReply, error) {
	out := new(UpfPfcpHeartbeatsSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfPfcpPolicerSet(ctx context.Context, in *UpfPfcpPolicerSet) (*UpfPfcpPolicerSetReply, error) {
	out := new(UpfPfcpPolicerSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfPfcpPolicerShow(ctx context.Context, in *UpfPfcpPolicerShow) (*UpfPfcpPolicerShowReply, error) {
	out := new(UpfPfcpPolicerShowReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpfPfcpReencode(ctx context.Context, in *UpfPfcpReencode) (*UpfPfcpReencodeReply, error) {
	out := new(UpfPfcpReencodeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfPfcpServerSet(ctx context.Context, in *UpfPfcpServerSet) (*UpfPfcpServerSetReply, error) {
	out := new(UpfPfcpServerSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfPfcpServerShow(ctx context.Context, in *UpfPfcpServerShow) (*UpfPfcpServerShowReply, error) {
	out := new(UpfPfcpServerShowReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpfPolicyAddDel(ctx context.Context, in *UpfPolicyAddDel) (*UpfPolicyAddDelReply, error) {
	out := new(UpfPolicyAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UpfPolicyDump(ctx context.Context, in *UpfPolicyDump) (RPCService_UpfPolicyDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_UpfPolicyDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_UpfPolicyDumpClient interface {
	Recv() (*UpfPolicyDetails, error)
	api.Stream
}

type serviceClient_UpfPolicyDumpClient struct {
	api.Stream
}

func (c *serviceClient_UpfPolicyDumpClient) Recv() (*UpfPolicyDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *UpfPolicyDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) UpfUpdateApp(ctx context.Context, in *UpfUpdateApp) (*UpfUpdateAppReply, error) {
	out := new(UpfUpdateAppReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
