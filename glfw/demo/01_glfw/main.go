package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"

	// "github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// 初始化GLFW
	err := glfw.Init()
	if err != nil {
		fmt.Println("glfw Init : ", err)
		return
	}

	defer glfw.Terminate()
	// 推荐用  "github.com/go-gl/glfw/v3.3/glfw"
	// opengl 主版本
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	// opengl 副版本号
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// opengl 模式
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	// 窗口是否可调整大小
	glfw.WindowHint(glfw.Resizable, glfw.False) // 不可调整
	// glfw.WindowHint(glfw.Resizable, glfw.True)

	// glfwWindowHint(GLFW_DECORATED, GL_FALSE);
	// glfw.WindowHint(glfw.Decorated, glfw.True)
	// 没有边框和标题栏
	glfw.WindowHint(glfw.Decorated, glfw.False)

	// 创建窗口
	window, err := glfw.CreateWindow(500, 500, "Demo1中文", nil, nil)

	//fullScreen := glfw.GetPrimaryMonitor()

	//window, err := glfw.CreateWindow(1920, 1080, "Demo1中文", fullScreen, nil)
	if err != nil {
		fmt.Println("glfw window : ", err)

		return
	}

	// 设置上下文
	window.MakeContextCurrent()

	// 初始化opengl
	if err := gl.Init(); err != nil {
		fmt.Println("gl.Init : ", err)
		return
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// 循环绘制函数
	for !window.ShouldClose() {

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		window.SwapBuffers()
		// 接收事件
		glfw.PollEvents()
	}

}
