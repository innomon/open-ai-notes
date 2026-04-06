# Numerical stability

[Source](https://share.google/aimode/QDCTiiOvwHHmz9FtU)

Numerical stability is critical when implementing Mish because the $e^x$ term in the underlying Softplus function can cause floating-point overflow for large positive inputs. In a stable implementation, we use an approximation for large $x$ values where $\text{softplus}(x) \approx x$. [1] 

## 1. Identify the Stability Threshold

The Softplus function $\ln(1 + e^x)$ behaves linearly for large values of $x$. For example, if $x > 20$, then $e^x$ is already over $4.8 \times 10^8$. At this point, $\ln(1 + e^x)$ is virtually indistinguishable from $x$. By switching to a linear identity function for large $x$, we prevent the computer from trying to calculate infinity. [2] 

## 2. Implement Conditional Logic in GoMLX

In GoMLX, you use `graph.Where` to create this conditional logic. This ensures the graph compiler (XLA) chooses the safe path during execution.

```go
func SafeMish(x *graph.Node) *graph.Node {
    g := x.Graph()
    dtype := x.DType()
    
    // Threshold for switching to linear approximation (typically 20.0)
    threshold := graph.Scalar(g, dtype, 20.0)
    isLarge := graph.Greater(x, threshold)

    // Path A: Standard Softplus for small/medium values
    one := graph.Scalar(g, dtype, 1.0)
    softplus := graph.Log(graph.Add(one, graph.Exp(x)))
    
    // Path B: Linear approximation for large values (Softplus(x) ≈ x)
    // We use Where to pick the stable value before Tanh
    stableSoftplus := graph.Where(isLarge, x, softplus)
    
    return graph.Mul(x, graph.Tanh(stableSoftplus))
}
```

## 3. Handle Underflow for Large Negative Values

For very large negative values (e.g., $x < -20$), $e^x$ becomes a very small number close to zero.

* **Softplus Behavior**: $\ln(1 + \text{tiny number}) \approx 0$.
* **Mish Behavior**: $x \cdot \tanh(0) = 0$.

Standard `graph.Exp(x)` handles this underflow gracefully by returning $0$, which is numerically stable and does not require extra `Where` logic. [2] 

## 4. Optimize the Final Computation

By using `stableSoftplus`, the $\tanh$ operation also stays within a safe range. Since $\tanh(x)$ approaches $1$ for large $x$, for $x > 20$, the entire Mish function simplifies to $f(x) \approx x \cdot 1 = x$. This prevents "exploding gradients" during backpropagation. [3] 

## Final Result

The numerically stable Mish implementation in GoMLX is:

$$f(x) = \begin{cases} x \cdot \tanh(\ln(1 + e^x)) & \text{if } x \leq 20 \\ x \cdot \tanh(x) \approx x & \text{if } x > 20 \end{cases}$$ 

This ensures your model can train without encountering NaN (Not a Number) errors during deep network optimization. [2] 

Would you like to see how this affects gradient flow compared to a standard ReLU implementation?

```python
import numpy as np

def softplus(x):
    return np.log(1 + np.exp(x))

def mish(x):
    return x * np.tanh(softplus(x))

# Test values
test_values = np.array([-100, -50, 0, 20, 50, 100])
print(f"Softplus: {softplus(test_values)}")
print(f"Mish: {mish(test_values)}")
```

---

[1] [https://eureka.patsnap.com](https://eureka.patsnap.com/report-how-to-implement-a-numerically-stable-sigmoid-function-for-deep-learning-models-code-tests#:~:text=Computing%20sigmoid%20functions%20in%20the%20logarithmic%20domain,without%20losing%20precision%20or%20causing%20numerical%20instability.)
[2] [https://arxiv.org](https://arxiv.org/pdf/2202.03493)
[3] [https://www.geeksforgeeks.org](https://www.geeksforgeeks.org/deep-learning/vanishing-and-exploding-gradients-problems-in-deep-learning/)
