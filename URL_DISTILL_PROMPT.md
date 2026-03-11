# URL Distillation Prompt

When provided with a URL, follow these steps to summarize the content and extract all relevant links:

1. **Information Gathering:**
   - Use `web_fetch` to retrieve the page title, main content, and any listed links directly.
   - If `web_fetch` returns generic footer content or fails, use `google_web_search` with the query: `"URL" page title summary and key links` to find metadata and content summaries from search snippets or secondary sources.

2. **Extraction Requirements:**
   - **Title:** The full, official title of the page or article.
   - **TL;DR:** A concise (3-5 bullet points) summary of the page's core message, key arguments, or main takeaways.
   - **Summary:** A structured summary of the page's content, focusing on key themes and information.
   - **Links:** Identify all relevant URLs mentioned in the content. Map them to their intended purpose (e.g., "Source Code," "Reference Article," "Referenced Tool").

3. **Formatting:**
   - Use the following Markdown structure:
     ```markdown
     ## [Page Title](URL)

     ### Content Summary
     [Brief summary of the page]

     ### Relevant Links Found in the Page
     - [Description of Link 1](URL 1)
     - [Description of Link 2](URL 2)
     ```

4. **Edge Cases:**
   - If a link is to a PDF or download, mark it as `(PDF)` or `(Download)`.
   - If a link is an affiliate link or product page, label it clearly (e.g., `[Product Name (Affiliate)]`).
   - If no specific description is found for a link, use the site name or a short snippet of surrounding text as the label.

5. **Save to File:**
   - Create an appropriate markdown filename: lower case, hyphen-separated with `.md` extension.
   - Example: `abc-def-xyz.md`
   - If the filename exists, append a timestamp: `abc-def-xyz-YYYYMMDD-HHMM.md`
