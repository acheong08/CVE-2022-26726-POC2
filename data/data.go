package c2Data

type Cmd struct{
        Id string
        Os string
        Target string
        Command string
        Kind string
}

type CmdOut struct{
        Id string
        Session string
        Output string
        Kind string
}
