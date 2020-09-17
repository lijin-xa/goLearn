module test

// repalce 引用其他包的 - 方法
replace add => ../add

go 1.14

require add v0.0.0-00010101000000-000000000000 // indirect
