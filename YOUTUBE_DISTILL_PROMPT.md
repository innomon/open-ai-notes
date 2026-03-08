# YouTube Distillation Prompt

When provided with a YouTube URL, follow these steps to summarize the content and extract all relevant links:

1. **Information Gathering:**
   - Use `web_fetch` to attempt to retrieve the video title, description, and any listed links directly.
   - If `web_fetch` returns generic footer content or fails, use `google_web_search` with the query: `"VIDEO_URL" video title description and links` to find metadata and description summaries from search snippets or secondary sources.

2. **Extraction Requirements:**
   - **Title:** The full, official title of the video.
   - **TL;DR:** A concise (3-5 bullet points) summary of the video's core message or key takeaways.
   - **Summary** summarize the video
   - **Links:** Identify every URL mentioned in the description. Map them to their intended purpose (e.g., "Source Code," "Article Mentioned," "Referenced Tool").

3. **Formatting:**
   - Use the following Markdown structure:
     ```markdown
     ## [Video Title](Video URL)

     ### Description Summary
     [Brief summary of the video]

     ### Links Referred in the Video and Description
     - [Description of Link 1](URL 1)
     - [Description of Link 2](URL 2)
     ```

4. **Edge Cases:**
   - If a link is a playlist, mark it as `(Playlist)`.
   - If a link is an affiliate link or product page, label it clearly (e.g., `[Product Name (Affiliate)]`).
   - If no specific description is found for a link, use the site name or a short snippet of surrounding text as the label.

5. **Save to File**
   - create an appropriate markdown filename : lower case hyper separated with `.md` extension
   - example `abc-def-xyz.md`
   - if the filename exist append a timestamp, like `abc-def-xyz-YYYYMMDD-HHMM.md`


