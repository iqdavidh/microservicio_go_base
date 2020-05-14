

14
##Acceso al body desde middleeare
TIL that http.Response.Body is a buffer, which means that once it has been read, it cannot be read again.

It's like a stream of water, you can see it and measure it as it passes but once it's gone, it's gone.

However, knowing this, there is a workaround, you need to "catch" the body and restore it:
```
// Read the Body content
var bodyBytes []byte
if context.Request().Body != nil {
    bodyBytes, _ = ioutil.ReadAll(context.Request().Body)
}

// Restore the io.ReadCloser to its original state
context.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

// Continue to use the Body, like Binding it to a struct:
order := new(models.GeaOrder)
error := context.Bind(order)
```
Now, you can use context.Request().Body somewhere else.

Sources:

---
Ejemplos del mtodos directamtente con gin
```
value, ok := db[nombre]
if ok {
    //ctx.JSON(http.StatusOK, libapi.DicJson{"nombre": nombre, "value": value})
    ctx.JSON(http.StatusOK, gin.H{"nombre": nombre, "value": value})
} else {
    ctx.JSON(http.StatusOK, gin.H{"nombre": nombre, "status": "no value"})
}
```
