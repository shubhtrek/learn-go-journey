# Step 4.2: Channels 📞

Channels are pipelines that connect concurrent goroutines. You send data into channels and receive it elsewhere.

## Key Concepts
1. **Declaration**: `ch := make(chan dataType)`
2. **Send**: `ch <- value`
3. **Receive**: `value := <-ch`
4. **Blocking**: Channel operations block until both sender and receiver are ready.

## How to Run
```bash
go run main.go
```
