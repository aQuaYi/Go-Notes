# exec模块

exec 是对os.StartProcess 的封装，方便执行外部命令。
## Cmd
```go
type Cmd
    // Command 生成一个待执行的*cmd
    func Command(name string, arg ...string) *Cmd
    // CommandContext 生成一个带context的待执行*cmd
    func CommandContext(ctx context.Context, name string, arg ...string) *Cmd
    // CombineOutput 执行命令，并合并 stdout 和 stderr 一起，返回输出结果。
    func (c *Cmd) CombinedOutput() ([]byte, error)
    // Output 执行命令，并返回 stdout 的输出结果。
    func (c *Cmd) Output() ([]byte, error)
    // Run 运行程序，并等待程序运行结束。阻塞。
    func (c *Cmd) Run() error
    // Start 启动一个程序，直接返回。异步。
    func (c *Cmd) Start() error
    // StderrPipe returns a pipe that will be connected to the command's standard error when the command starts
    func (c *Cmd) StderrPipe() (io.ReadCloser, error)
    // StdinPipe returns a pipe that will be connected to the command's standard input when the command starts. 
    func (c *Cmd) StdinPipe() (io.WriteCloser, error)
    // StdoutPipe returns a pipe that will be connected to the command's standard output when the command starts.
    func (c *Cmd) StdoutPipe() (io.ReadCloser, error)

    // Wait waits for the command to exit. It must have been started by Start.
    // Wait 会释放cmd所有的相关资源，这样也就省略了stderrPipe和stdoutPipe的显性关闭
    func (c *Cmd) Wait() error
```
## Run和Start的区别

Run = Start + Wait

## StderrPipe和StdoutPipe
一定要等到读取完成后，再使用Wait查看结果。