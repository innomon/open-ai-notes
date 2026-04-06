# Mish GoMLX training loop

[Source](https://share.google/aimode/hYXqfCgSFox2R1CxJ)

In [GoMLX](https://github.com/gomlx/gomlx), you incorporate a custom activation function like Mish by defining it as a `graph.Node` transformation and applying it within your `modelGraph` function. GoMLX uses automatic differentiation, so you only need to define the forward pass. [1, 2] 

## 1. Define the Mish Function [3] 

This function takes an input node and returns a new node with the Mish transformation applied. [2] 

```go
import (
    "github.com/gomlx/gomlx/graph"
)

// Mish: x * Tanh(Softplus(x))
func Mish(x *graph.Node) *graph.Node {
    g := x.Graph()
    dtype := x.DType()
    
    // Softplus(x) = Log(1 + Exp(x))
    one := graph.Scalar(g, dtype, 1.0)
    softplus := graph.Log(graph.Add(one, graph.Exp(x)))
    
    return graph.Mul(x, graph.Tanh(softplus))
}
```

## 2. Integrate into the Model Graph

Use the custom Mish function after a layer (like `layers.Dense`) instead of a built-in activation. [2] 

```go
func modelGraph(ctx *context.Context, spec any, inputs []*graph.Node) []*graph.Node {
    x := inputs[0]
    
    // Layer 1: Dense -> Mish
    x = layers.Dense(ctx.In("layer1"), x, true, 128)
    x = Mish(x)
    
    // Layer 2: Dense -> Softmax (Output)
    logits := layers.Dense(ctx.In("output"), x, true, 10)
    return []*graph.Node{logits}
}
```

## 3. Setup the Training Loop

The training loop in [GoMLX](https://pkg.go.dev/github.com/gomlx/gomlx/pkg/ml/context) uses a `Trainer` object and a `Dataset`. [2, 4] 

```go
func Train() {
    // 1. Initialize context and backend (e.g., XLA)
    ctx := context.NewContext(backend)
    
    // 2. Define Loss and Optimizer
    lossFn := func(ctx *context.Context, spec any, inputs, outputs []*graph.Node) *graph.Node {
        return losses.SoftmaxCrossEntropy(outputs[0], inputs[1])
    }
    optimizer := optimizers.Adam().LearningRate(0.001).Done()
    
    // 3. Create Trainer
    trainer := train.NewTrainer(backend, ctx, modelGraph, lossFn, optimizer)
    
    // 4. Run loop
    for epoch := 0; epoch < 50; epoch++ {
        dataset.Reset()
        for dataset.Next() {
            trainer.TrainStep(dataset.Inputs(), dataset.Labels())
        }
        fmt.Printf("Epoch %d: Loss = %v\n", epoch, trainer.LastLoss())
    }
}
```

## Summary of Workflow

* **Automatic Differentiation**: GoMLX automatically handles the complex derivative of Mish ($x \cdot \tanh(\ln(1+e^x))$) during the `TrainStep`.
* **Context Scoping**: Use `ctx.In("name")` to ensure weights for each layer are managed correctly in the global context.
* **Flexibility**: Unlike rigid frameworks, you can swap Mish for ReLU or Swish simply by changing the function call in `modelGraph`.

Would you like to explore how to add numerical stability to this Mish implementation for high-precision training?

---

[1] [https://arxiv.org](https://arxiv.org/abs/1908.08681)
[2] [https://levelup.gitconnected.com](https://levelup.gitconnected.com/machine-learning-in-go-part-1-f8ef8f489c91)
[3] [https://www.youtube.com](https://www.youtube.com/watch?v=gH3XpUYWilY&t=235)
[4] [https://keras.io](https://keras.io/guides/writing_a_custom_training_loop_in_tensorflow/)
