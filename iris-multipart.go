func save(ctx *iris.Context)  {
    // Get name and email
    name := ctx.FormValueString("name")
    email := ctx.FormValueString("email")
    // Get avatar
    avatar, err := ctx.FormFile("avatar")
    if err != nil {
       ctx.EmitError(iris.StatusInternalServerError)
       return
    }

    // Source
    src, err := avatar.Open()
    if err != nil {
       ctx.EmitError(iris.StatusInternalServerError)
       return
    }
    defer src.Close()

    // Destination
    dst, err := os.Create(avatar.Filename)
    if err != nil {
       ctx.EmitError(iris.StatusInternalServerError)
       return
    }
    defer dst.Close()

    // Copy
    if _, err = io.Copy(dst, src); err != nil {
       ctx.EmitError(iris.StatusInternalServerError)
       return
    }

    ctx.HTML(iris.StatusOK, "<b>Thanks!</b>")
}
