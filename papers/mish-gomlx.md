# Mish activation function

[Source](https://share.google/aimode/OhnMaMdYGu3KfUkqc)

Implementing the Mish activation function in the [GoMLX framework](https://github.com/gomlx/gomlx) is straightforward because GoMLX provides a set of differentiable operators that allow you to define custom functions easily. [1, 2] 

Mish is mathematically defined as:

$$f(x) = x \cdot \tanh(\text{softplus}(x)) = x \cdot \tanh(\ln(1 + e^x))$$ 

## Implementation Code

You can implement this in Go using the `github.com/gomlx/gomlx` package. Since GoMLX does not currently include a built-in Mish type in its standard activations enum, you define it as a custom function that operates on `*graph.Node` objects: [3, 4] 

```go
import (
    "github.com/gomlx/gomlx/graph"
)

// Mish activation function: x * Tanh(Softplus(x))
func Mish(x *graph.Node) *graph.Node {
    // Softplus(x) = Log(1 + Exp(x))
    softplus := graph.Log(graph.Add(graph.Scalar(x.Graph(), x.DType(), 1.0), graph.Exp(x)))
    
    // Mish = x * Tanh(softplus)
    return graph.Mul(x, graph.Tanh(softplus))
}
```

## How to use it in a Model

Once defined, you can manually apply this function to your layers instead of using the standard `activations.Apply` method.

* **Manual Application**: Use it directly on the output of a dense or convolutional layer.

```go
// Example: Dense layer followed by Mish
logits := layers.Dense(ctx, input, units)
activated := Mish(logits)
```

* **Numerical Stability (Optional)**: For very large values of $x$, `Exp(x)` can overflow. In many frameworks, Softplus is implemented as $x$ when $x$ is large (e.g., $x > 20$). You can add a `graph.Where` condition if you encounter stability issues, though GoMLX's underlying XLA/OpenXLA engine often handles this during JIT compilation. [2, 5] 

## Why use Mish in GoMLX?

* **No Manual Derivatives**: Because GoMLX uses automatic differentiation, you only need to define the "forward" pass (the formula itself). The framework will automatically calculate the gradients for backpropagation.
* **Performance**: Mish is non-monotonic and smooth, which often leads to better accuracy in deep models (like YOLOv4) compared to ReLU. [5, 6, 7, 8] 

Would you like to see how to incorporate this custom activation into a full GoMLX training loop?

---

[1] [https://pkg.go.dev](https://pkg.go.dev/github.com/gomlx/gomlx/pkg/ml/layers/activations)
[2] [https://github.com](https://github.com/gomlx/gomlx)
[3] [https://pkg.go.dev](https://pkg.go.dev/github.com/gomlx/gomlx/pkg/ml/layers/activations#:~:text=Type%20is%20an%20enum%20for%20the%20supported,to%20TypeSwish%20TypeSilu%20TypeTanh%20TypeGelu%20TypeGeluApprox%20%29)
[4] [https://pkg.go.dev](https://pkg.go.dev/github.com/gomlx/gomlx/pkg/ml/layers/activations#:~:text=Links%20*%20Apply%28activation%2C%20x%29%20ApplyFromContext%28ctx%2C%20x%29%20Gelu%28x%29,%28i%29%20UnmarshalJSON%28data%29%20%28i%29%20UnmarshalText%28text%29%20%28i%29%20UnmarshalYAML%28unmarshal%29%20Values%28%29)
[5] [https://www.youtube.com](https://www.youtube.com/watch?v=rjh8ju883cQ&t=130)
[6] [https://www.youtube.com](https://www.youtube.com/watch?v=4eIlvlP-8wk)
[7] [https://arxiv.org](https://arxiv.org/abs/1908.08681)
[8] [https://medium.com](https://medium.com/@shravankoninti/build-a-small-language-model-slm-from-scratch-3ddd13fa6470#:~:text=forward%28self%2C%20idx%2C%20targets=None%29%20%2D%20The%20Forward%20Pass,when%20you%20pass%20data%20through%20the%20model.)
