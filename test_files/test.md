# GitHub Markdown Test File

This file includes all GitHub Flavored Markdown formatting options and their cross variants for testing purposes.

## Headings

# H1 Heading
## H2 Heading
### H3 Heading
#### H4 Heading
##### H5 Heading
###### H6 Heading

## Styling Text

**Bold text**
__Also bold__
*Italic text*
_Also italic_
~~Strikethrough~~
<sub>Subscript</sub>
<sup>Superscript</sup>
<ins>Underline</ins>

### Cross Variants
**Bold and _italic_ combined**
***All bold and italic***
~~Strikethrough with **bold** and *italic*~~
<sub>Subscript with <sup>superscript</sup></sub>

## Quoting Text

Normal text.

> Quoted text.
>
> Multiple lines in quote.

### Nested Quotes
> First level quote
>> Second level quote

## Quoting Code

Inline code: `git status`

Code block:
```
git status
git add .
git commit -m "message"
```

Syntax highlighted code block:
```javascript
function hello() {
  console.log("Hello, world!");
}
```

## Supported Color Models

Colors: `#FF0000` `rgb(255,0,0)` `hsl(0,100%,50%)`

## Links

[Inline link](https://github.com)

[Link with title](https://github.com "GitHub")

Autolink: https://github.com

## Section Links

[Link to H1](#h1-heading)

## Relative Links

[Relative link](./test.md)

## Custom Anchors

<a name="custom-anchor"></a>

[Link to custom anchor](#custom-anchor)

## Line Breaks

Line 1  
Line 2

Line 1\
Line 2

Line 1<br/>
Line 2

Paragraph 1

Paragraph 2

## Images

![Alt text](https://github.com/images/error/octocat_happy.gif)

## Lists

### Unordered Lists
- Item 1
- Item 2
  - Nested item
  - Another nested

### Ordered Lists
1. First item
2. Second item
   1. Nested ordered
   2. Another nested

### Mixed Lists
1. Ordered
   - Unordered under ordered
2. Back to ordered

## Task Lists

- [x] Completed task
- [ ] Incomplete task
- [x] Task with **bold** and *italic*
- [ ] ~~Strikethrough task~~

## Emojis

:smile: :+1: :tada:

## Paragraphs

This is a paragraph.

This is another paragraph.

## Footnotes

Here is a footnote[^1].

Another footnote[^2].

[^1]: First footnote.
[^2]: Second footnote with  
     multiple lines.

## Alerts

> [!NOTE]
> This is a note.

> [!TIP]
> This is a tip.

> [!IMPORTANT]
> This is important.

> [!WARNING]
> This is a warning.

> [!CAUTION]
> This is a caution.

## Hiding Content with Comments

<!-- This content is hidden -->

Visible content.

## Ignoring Markdown Formatting

\*Not italic\*

## Tables (Advanced, but included)

| Header 1 | Header 2 |
|----------|----------|
| Cell 1   | Cell 2   |
| Cell 3   | Cell 4   |

## Cross Variants Examples

### Nested Everything
1. **Bold item**
   - *Italic subitem*
     - ~~Strikethrough subsubitem~~
       - `Code in list`

### Code with Links and Emphasis
```markdown
[Link in code](url) *not italic*
```

### Quotes with Formatting
> **Bold quote** with `code` and [link](url)

### Footnotes with Formatting
Footnote with **bold**[^3].

[^3]: Footnote with *italic* and `code`.

### Alerts with Cross Elements
> [!WARNING]
> Warning with **bold**, *italic*, `code`, and [link](url).

## End of Test File