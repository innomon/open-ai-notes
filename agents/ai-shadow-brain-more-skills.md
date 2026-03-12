# AI Shadow Brain: More Skills (no .md)

## [AI Shadow Brain: More Skills (no .md)](https://www.youtube.com/watch?v=XCQPjaIAeeM)

### TL;DR
- **Catastrophic Forgetting Fixed:** Introduces a breakthrough method from Google Research and University of Wisconsin-Madison to prevent LLMs from losing old knowledge while learning new skills.
- **"Grow, Don't Overwrite" (GDO):** A technique that expands the Multi-Layer Perceptron (MLP) dimensions of a model instead of overwriting existing weights.
- **Function Preservation:** New parameters are initialized to ensure the model's output remains identical at the start of training, preserving original capabilities.
- **Permanent Core Integration:** Moves beyond external markdown-based "skills" to integrated, permanent knowledge within the LLM's architecture.

### Description Summary
The video discusses a significant advancement in AI training called "Grow, Don't Overwrite" (GDO). Traditionally, fine-tuning an LLM on new specialized data (like deep mathematics or a new language) often leads to "catastrophic forgetting," where the model loses its original world knowledge. The GDO method avoids this by expanding the model's architecture. By freezing the original "memory" weights (the down projection) and only training the new "upper projection" and routing logic, the model can acquire complex new skills permanently. This "Shadow Brain" approach represents a shift towards AI agents that learn organically and continuously without degrading their base intelligence.

### Links Referred in the Video and Description
- [Grow, Don't Overwrite: Fine-tuning Without Forgetting (ArXiv Paper)](https://arxiv.org/abs/2603.08647) - The primary research paper discussed in the video.
- [Agentic Critical Training (ACT) (ArXiv Paper)](https://arxiv.org/abs/2603.07887) - Related research mentioned regarding agent training.
- [FlashAttention-4 (ArXiv Paper)](https://arxiv.org/abs/2603.05123) - Technical reference to memory-efficient attention mentioned in the context of scaling.
- [Google Research](https://research.google/) - Institutional source for the GDO research.
- [University of Wisconsin-Madison](https://www.wisc.edu/) - Academic collaborator for the research paper.
