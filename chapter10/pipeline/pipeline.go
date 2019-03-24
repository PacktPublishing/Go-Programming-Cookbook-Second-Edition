package pipeline

import "context"

// NewPipeline initializes the workers and
// connects them, it returns the input of the pipeline
// and the final output
func NewPipeline(ctx context.Context, numEncoders, numPrinters int) (chan string, chan string) {
	inEncode := make(chan string, numEncoders)
	inPrint := make(chan string, numPrinters)
	outPrint := make(chan string, numPrinters)
	for i := 0; i < numEncoders; i++ {
		w := Worker{
			in:  inEncode,
			out: inPrint,
		}
		go w.Work(ctx, Encode)
	}

	for i := 0; i < numPrinters; i++ {
		w := Worker{
			in:  inPrint,
			out: outPrint,
		}
		go w.Work(ctx, Print)
	}
	return inEncode, outPrint
}
