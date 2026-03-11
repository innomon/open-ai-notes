# YouTube Distillation Prompt

When provided with a YouTube URL, follow these steps to summarize the content and extract all relevant links:

1. **Extraction and Verification:**
   - Extract the unique Video ID (e.g., `abc123XYZ00`) from the provided URL.
   - Use `web_fetch` to retrieve the video title, description, and links.
   - **Verification:** Ensure the retrieved metadata matches the specific Video ID. If `web_fetch` returns a generic page, a search result list, or a "Video Unavailable" message, proceed to Step 2.

2. **Targeted Search Fallback:**
   - If `web_fetch` fails, use `google_web_search` with a strict query: `site:youtube.com "VIDEO_URL"` or `site:youtube.com "Video ID" description links`.
   - **Crucial Warning:** Avoid searching for descriptive phrases like "how to find description" or "where are the links," as these often lead to generic help videos (e.g., Richard Rost's "Where Are The Links?"). If the search results do not clearly show the title and description for the *exact* video requested, report that the information could not be retrieved.

3. **Extraction Requirements:**
   - **Title:** The full, official title of the video.
   - **TL;DR:** A concise (3-5 bullet points) summary of the video's core message or key takeaways.
   - **Summary:** A structured summary of the video's content.
   - **Links:** Identify every URL mentioned in the description. Map them to their intended purpose (e.g., "Source Code," "Article Mentioned," "Referenced Tool").

4. **Formatting:**
   - Use the following Markdown structure:
     ```markdown
     ## [Video Title](Video URL)

     ### Description Summary
     [Brief summary of the video]

     ### Links Referred in the Video and Description
     - [Description of Link 1](URL 1)
     - [Description of Link 2](URL 2)
     ```

5. **Edge Cases:**
   - If a link is a playlist, mark it as `(Playlist)`.
   - If a link is an affiliate link or product page, label it clearly (e.g., `[Product Name (Affiliate)]`).
   - If no specific description is found for a link, use the site name or a short snippet of surrounding text as the label.

6. **Save to File:**
   - Create an appropriate markdown filename: lower case, hyphen-separated with `.md` extension.
   - Example: `abc-def-xyz.md`
   - If the filename exists, append a timestamp: `abc-def-xyz-YYYYMMDD-HHMM.md`


