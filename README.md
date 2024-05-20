# PUZZLE-CAPTCHA

一个简易的滑动图形验证码图片处理。

## USAGE

```golang
f, _ := os.Open("image.png")
img, _ := png.Decode(f)
c := NewPuzzleCaptcha(img, 50, 2)
base, hole := c.Generate()
```
