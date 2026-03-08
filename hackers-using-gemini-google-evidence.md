## [Hackers Are Using Gemini — Google Published the Evidence](https://youtu.be/gKJW1OpFuCk?si=SgZrXS7Antie0fBI)

### Description Summary
This video by Claudius analyzes a major report from Google's Threat Intelligence Group (GTIG) revealing how state-sponsored hacking groups (from Iran, China, North Korea, and Russia) and cybercriminals are integrating Gemini and other AI tools into their operations. It details the shift from using AI as a simple coding assistant to dynamic, "agentic" use cases where malware queries LLMs in real-time to evade detection and automate attacks.

### TL;DR
- **State-Sponsored Usage:** Government-backed hacking groups from Iran, China, North Korea, and Russia are actively using Google's Gemini and other AI tools for cyber operations.
- **On-the-Fly Malware:** New malware strains (like HONESTCUE) call LLM APIs during an attack to generate payloads in memory, bypassing disk-based antivirus detection.
- **Self-Mutating Code:** Tools like PROMPTFLUX use AI to rewrite their own code hourly, making them extremely difficult to signature and track.
- **Social Engineering at Scale:** AI is being used by groups like APT42 to build rapport and create more convincing phishing campaigns tailored to specific targets.

### Summary
The video explores the findings of Google's GTIG regarding the adversarial use of AI. It highlights how actors have moved beyond using LLMs for brainstorming to integrating them directly into the "kill chain." Notable examples include APT42 using Gemini for "rapport-building" in phishing, and APT28 being the first to have malware query an LLM during live operations. A significant technological leap is seen in malware that uses APIs to generate malicious code on-the-fly, ensuring that the malicious payload never touches the disk, which traditional security tools monitor. The video concludes that as AI agents become more prevalent, the defense against these automated, intelligent attacks will require a fundamental shift in cybersecurity strategy.

### Links Referred in the Video and Description
- [Google GTIG AI Threat Tracker (February 2026)](https://cloud.google.com/blog/topics/threat-intelligence/distillation-experimentation-integration-ai-adversarial-use) - Official report on AI distillation and adversarial integration.
- [Google GTIG AI Threat Tracker (November 2025)](https://cloud.google.com/blog/topics/threat-intelligence/threat-actor-usage-of-ai-tools) - Previous report on AI tool usage by threat actors.
- [Malwarebytes — Fake Gemini chatbot / Google Coin scam](https://www.malwarebytes.com/blog/ai/2026/02/scammers-use-fake-gemini-ai-chatbot-to-sell-fake-google-coin) - Case study of a fake Gemini chatbot selling a non-existent cryptocurrency.
- [MIT Technology Review — Cyberattacks by AI Agents](https://www.technologyreview.com/2025/04/04/1114228/cyberattacks-by-ai-agents-are-coming/) - Expert commentary on the rise of agentic AI in cyber warfare.
