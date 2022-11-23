# eInk Radiator Image Source: Text

![CI](https://ci.petewall.net/api/v1/teams/main/pipelines/eink-radiator/jobs/test-image-source-text/badge)

Generates an image from text.

```bash
text generate --config config.json --height 300 --width 400
```

## Configuration

The configuration describes the text to render and how to render it

| field            | default | required | description |
|------------------|---------|----------|-------------|
| text             |         | Yes      | The text to render |
| wrap             | false   | No       | Automatically add line breaks to fit the text to the width of the image |
| font             | Ubuntu  | No       | The font to use |
| size             | 0       | No       | The font size to use (0 will resize to fit to the size of the image) |
| color            | black   | No       | The color of the text |
| background.color | white   | No       | The color of the background |

## Examples

All examples are rendered to 400 x 300 resolution

### Text alone

Renders the text, resizing to fit the screen

```yaml
---
text: One Two Three Four Five Six Seven Eight Nine Ten
```

![An image rendering the simple text example](test/outputs/text_fit.png)

### All fields

```yaml
text: |-
  Shields up!
  Rrrrrred Alert!
color: red
background:
  color: black
font: UbuntuMono
size: 48
```

![An image rendering the full example](test/outputs/colors.png)
