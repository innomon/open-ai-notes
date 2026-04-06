# Comparison SoftPlus with Swish and Mish

[https://share.google/aimode/ldjkxez6yG0ST2tPy](https://share.google/aimode/ldjkxez6yG0ST2tPy)

While Softplus is a smooth, monotonic approximation of ReLU, Swish and Mish are more modern, non-monotonic functions designed specifically to improve information flow in deep networks. [1, 2] 

## Comparison Table

| Feature [3, 4, 5, 6, 7] | Softplus | Swish | Mish |
|---|---|---|---|
| Formula | $\ln(1 + e^x)$ | $x \cdot \sigma(\beta x)$ | $x \cdot \tanh(\text{softplus}(x))$ |
| Monotonicity | Monotonic (always increases) | Non-monotonic | Non-monotonic |
| Lower Bound | $0$ | $\approx -0.28$ (for $\beta=1$) | $\approx -0.31$ |
| Key Advantage | Smooth everywhere; strictly positive. | Self-gated; better generalization than ReLU. | Superior gradient flow in very deep networks. |
| Best Use Case | When strictly positive outputs are required. | General deep learning; mobile/efficient nets. | State-of-the-art accuracy in complex architectures. |

---

## 1. Softplus: The Smooth Baseline
Softplus acts as a "soft" version of ReLU. It is monotonic, meaning its output never decreases as the input increases. While it avoids the sharp discontinuity of ReLU at zero, it is often outperformed by Swish and Mish in modern deep learning tasks because it does not allow for negative activations. [1, 7, 8, 9, 10] 

## 2. Swish: The Self-Gated Alternative
Developed by Google researchers, Swish is defined as $x$ multiplied by the Sigmoid function. [11, 12, 13, 14] 

* **Non-monotonicity**: Unlike Softplus, Swish has a small "dip" below zero for negative inputs. This property helps maintain small negative weights, which can be crucial for learning complex patterns.
* **Performance**: It has been shown to improve accuracy on ImageNet by approximately 0.6% over ReLU for certain architectures. [5, 6, 12] 

## 3. Mish: The High-Performer
Mish is a self-regularized function that combines identity, tanh, and softplus. [15, 16] 

* **Increased Smoothness**: Mish is smoother than both Swish and ReLU, which helps the gradient flow more effectively through many layers.
* **Superior Accuracy**: In many benchmarks (like YOLOv4 and ResNet-50), Mish consistently outperforms Swish and ReLU, often providing a 1% to 3% improvement in accuracy.
* **Computational Cost**: Its main drawback is being more mathematically complex, making it slightly slower to compute than Swish or ReLU. [5, 17, 18, 19, 20] 

## Summary of Differences

* Softplus is a simple, smooth version of ReLU but lacks the advanced "self-gating" benefits of the others.
* Swish is efficient and robust, frequently used as a drop-in replacement for ReLU.
* Mish is currently considered the most powerful of the three for deep networks, though it comes with a higher computational price tag. [1, 6, 20, 21] 

---

### References
[1] [https://www.johndcook.com](https://www.johndcook.com/blog/2023/08/06/swish-mish-and-serf/)  
[2] [https://www.youtube.com](https://www.youtube.com/watch?v=WXl-lQK5WHY#:~:text=Softplus%20is%20smooth%20&%20differentiable%20at%200%2C,ReLU%2C%20&%20Swish%20is%20self%2Dgated%20&%20non%2Dmonotonic.)  
[3] [https://www.geeksforgeeks.org](https://www.geeksforgeeks.org/machine-learning/activation-functions-neural-networks/)  
[4] [https://machine-learning-made-simple.medium.com](https://machine-learning-made-simple.medium.com/how-google-discovered-the-swish-and-many-other-activation-function-9639130a6ad3)  
[5] [https://arxiv.org](https://arxiv.org/pdf/1908.08681)  
[6] [https://amitnikhade.medium.com](https://amitnikhade.medium.com/swish-activation-function-d106fe13930e)  
[7] [https://www.linkedin.com](https://www.linkedin.com/posts/chiragsubramanian_the-revolutionary-mish-activation-function-activity-7317792022466105344-__5s)  
[8] [https://www.researchgate.net](https://www.researchgate.net/publication/400188096_Rethinking_Differentiability_A_Comparative_Study_of_Smooth_vs_Non-Smooth_Activation_Functions_in_CNNs)  
[9] [https://www.johndcook.com](https://www.johndcook.com/blog/2023/08/06/swish-mish-and-serf/)  
[10] [https://dergipark.org.tr](https://dergipark.org.tr/en/download/article-file/4895272)  
[11] [https://www.geeksforgeeks.org](https://www.geeksforgeeks.org/python/swish-activation-function-in-pytorch/)  
[12] [https://www.geeksforgeeks.org](https://www.geeksforgeeks.org/machine-learning/ml-swish-function-by-google-in-keras/)  
[13] [https://www.geeksforgeeks.org](https://www.geeksforgeeks.org/deep-learning/swish-activation-function/#:~:text=S%20w%20i%20s%20h%20%28%20x,ImageNet%20dataset%20by%200.6%25%20for%20Inception%2DResNet%2Dv2%20architecture.)  
[14] [https://cognitivecreator.medium.com](https://cognitivecreator.medium.com/swish-activation-function-86b9d5865c9f)  
[15] [https://sefiks.com](https://sefiks.com/2019/10/28/mish-as-neural-networks-activation-function/)  
[16] [https://sh-tsang.medium.com](https://sh-tsang.medium.com/review-mish-a-self-regularized-non-monotonic-activation-function-a7afe19b4af7)  
[17] [https://sh-tsang.medium.com](https://sh-tsang.medium.com/review-mish-a-self-regularized-non-monotonic-activation-function-a7afe19b4af7)  
[18] [https://arxiv.org](https://arxiv.org/pdf/1908.08681)  
[19] [https://arxiv.org](https://arxiv.org/abs/1908.08681)  
[20] [https://www.mdpi.com](https://www.mdpi.com/2079-9292/11/4/540)  
[21] [https://lessw.medium.com](https://lessw.medium.com/meet-mish-new-state-of-the-art-ai-activation-function-the-successor-to-relu-846a6d93471f)
