package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// 替换颜色函数
func replaceColor(img image.Image, targetColor color.Color, replacementColor color.Color) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImage := image.NewRGBA(bounds)
	fmt.Println(targetColor)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// 获取当前像素颜色
			originalColor := img.At(x, y)
			_, _, _, a := originalColor.RGBA()
			// 判断颜色是否匹配
			if a != 0 {
				// fmt.Println(r/256, g/256, b/256, a/256)
				newImage.Set(x, y, replacementColor)
			} else {
				newImage.Set(x, y, originalColor)
			}
		}
	}

	return newImage
}

func main() {
	// 打开源图片文件
	file, err := os.Open("DayZ-logo-white.png")
	if err != nil {
		fmt.Println("Error opening image:", err)
		return
	}
	defer file.Close()

	// 解码PNG图片
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// 定义目标颜色和替换颜色（这里的颜色值是示例，可以根据需要修改）
	targetColor := color.RGBA{R: 35, G: 31, B: 32, A: 255}      // 红色, 31, 32
	replacementColor := color.RGBA{R: 35, G: 31, B: 32, A: 255} // 绿色

	// 替换颜色
	newImg := replaceColor(img, targetColor, replacementColor)

	// 保存修改后的图片
	outFile, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Error creating output image:", err)
		return
	}
	defer outFile.Close()

	// 将修改后的图片保存为PNG格式
	err = png.Encode(outFile, newImg)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	fmt.Println("Image processed and saved as output.png")
}
