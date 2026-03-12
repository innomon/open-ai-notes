## [Karpathy's Autoresearch: We Achieved Near-Human Scores in 2 Hours!](https://www.youtube.com/watch?v=9jxrmk_Xses)

### TL;DR
- **AI-Driven Research:** Explores Andrej Karpathy's `autoresearch` tool, which uses AI agents (like Claude Code) to automate the entire ML research workflow.
- **Consumer Hardware Results:** Achieved near-human prediction scores (0.511 val_bpb) on the TinyStories dataset using an RTX 3060 in under 2 hours.
- **Autonomous Optimization:** The AI agent conducted 33 autonomous experiments, discovering that shallower models and specific batch sizes were more effective than deeper architectures for this task.
- **Performance Jump:** Validation bits per byte (val_bpb) dropped 56% (from 1.17 to 0.511), nearly matching the human baseline of ~0.5.

### Description Summary
This video demonstrates how to use Andrej Karpathy's open-source `autoresearch` tool to automate machine learning experiments. Instead of manual parameter tweaking, a research plan is written in plain English, and the Claude Code AI agent executes multiple trials, analyzes results, and iterates autonomously. The demonstration focuses on the TinyStories dataset, showing how the agent optimized model parameters to move from near-gibberish output to coherent, human-level storytelling.

### Links Referred in the Video and Description
- [Tonbi (X/Twitter)](https://x.com/tonbistudio) - Creator's social profile
- [Tonbi's GitHub](https://github.com/tonbistudio) - Creator's GitHub repository
- [Portfolio](https://www.tonbistudio.com) - Creator's personal website
- [Karpathy's Autoresearch](https://github.com/karpathy/autoresearch) - Source code for the original tool
- [Windows Fork](https://github.com/jseam2/autoresearch) - Windows-compatible version of the tool
- [Claude Code](https://claude.ai/code) - AI coding assistant used for the research agent
- [TinyStories Dataset](https://huggingface.co/datasets/roneneldan/TinyStories) - Dataset used for training and testing
