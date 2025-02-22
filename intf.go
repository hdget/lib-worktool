package worktool

type WorkTool interface {
	SendFile(targets []string, file File) error
	SendText(targets []string, content string, atPeoples ...string) error
}
