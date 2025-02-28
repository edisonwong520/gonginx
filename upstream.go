package gonginx

import (
	"errors"
)

//Upstream represents `upstream{}` block
type Upstream struct {
	UpstreamName    string
	UpstreamServers []*UpstreamServer
	//Directives Other directives in upstream (ip_hash; etc)
	Directives []IDirective
}

//GetName Statement interface
func (us *Upstream) GetName() string {
	return "upstream"
}

//GetParameters upsrema parameters
func (us *Upstream) GetParameters() []string {
	return []string{us.UpstreamName} //the only parameter for an upstream is its name
}

//GetBlock upstream does not have block
func (us *Upstream) GetBlock() IBlock {
	return us
}

//GetDirectives get sub directives of upstream
func (us *Upstream) GetDirectives() []IDirective {
	directives := make([]IDirective, 0)
	directives = append(directives, us.Directives...)
	for _, uss := range us.UpstreamServers {
		directives = append(directives, uss)
	}

	return directives
}

func NewLuaBlock(directive IDirective) (u *LuaBlock, err error) {
	u = &LuaBlock{Name: directive.GetName()}
	return
}

//NewUpstream create new upstream from a directive
func NewUpstream(directive IDirective) (*Upstream, error) {
	parameters := directive.GetParameters()
	us := &Upstream{
		UpstreamName: parameters[0], //first parameter of the directive is the upstream name
	}

	if directive.GetBlock() == nil {
		return nil, errors.New("missing upstream block")
	}

	if len(directive.GetBlock().GetDirectives()) > 0 {
		for _, d := range directive.GetBlock().GetDirectives() {
			if d.GetName() == "server" {
				us.UpstreamServers = append(us.UpstreamServers, NewUpstreamServer(d))
			} else {
				us.Directives = append(us.Directives, d)
			}
		}
	}

	return us, nil
}

//AddServer add a server to upstream
func (us *Upstream) AddServer(server *UpstreamServer) {
	us.UpstreamServers = append(us.UpstreamServers, server)
}

//FindDirectives find directives in block recursively
func (us *Upstream) FindDirectives(directiveName string) []IDirective {
	directives := make([]IDirective, 0)
	for _, directive := range us.Directives {
		if directive.GetName() == directiveName {
			directives = append(directives, directive)
		}
		if directive.GetBlock() != nil {
			directives = append(directives, directive.GetBlock().FindDirectives(directiveName)...)
		}
	}

	return directives
}
