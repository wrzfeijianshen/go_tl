#include <iostream>
// #pragma comment(linker,"/subsystem:\"windows\" /entry:\"mainCRTStartup\"")

using namespace std;

int main(int argc,char** argv)
{
    //  -lstdc++   -mwindows
    // gcc -lstdc++  xxx.c 一般gcc编译cpp文件会这样用
    // -mwindows 是代表去掉黑窗口,而咱们需要这个黑窗口,所不能去掉
	// g++ -o bash.exe main.cpp  --> 生成的文件 C:\Program Files\Git\bin\bash.exe 即可

    // 调用系统命令即可. 
    system("bash");
    return 0;
}