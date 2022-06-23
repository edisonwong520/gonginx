package gonginx

type LuaBlock struct {
	Name string
}

func (s *LuaBlock) GetName() string {
	return "lua_block"
}

func (s *LuaBlock) GetParameters() (rsp []string) {
	return
}

func (s *LuaBlock) GetBlock() (rsp IBlock) {
	return
}
