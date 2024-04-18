// 这里包含了一些常量

package ss

// modifier 修饰符
type modifier uint8

const (
	Var   = modifier(1) //预留，但不支持
	Let   = modifier(2)
	Const = modifier(3)
)

// Operate 操作符 20
type Operate uint8

var (
	// 单目运算符
	unary = []Operate{Not, UnaryMinus, UnaryPlus, BitNot}
)

// 运算符
const (
	/** 算术运算符**/

	Plus     = Operate(10) // '+';
	Minus    = Operate(11) //'-';
	Multiply = Operate(12) //'*';
	Divide   = Operate(13) //'/';
	Modulus  = Operate(14) //'%';
	Power    = Operate(15) //'**';

	/** 比较运算符**/

	Equals            = Operate(16) //'==';
	IdentityEquals    = Operate(17) //'===';
	NotEquals         = Operate(18) // '!=';
	IdentityNotEquals = Operate(19) //'!==';
	MoreThan          = Operate(20) //'>';
	LessThan          = Operate(21) //'<';
	GreaterThanEquals = Operate(22) //'>=';
	LessThanEquals    = Operate(23) //'<=';

	/** 逻辑运算符**/

	And = Operate(24) //'&&';
	Or  = Operate(25) //'||';
	Not = Operate(26) //'!';

	/** 位运算符**/

	BitAnd               = Operate(27) //'&';
	BitOr                = Operate(28) //'|';
	BitXOr               = Operate(29) //'^';
	BitNot               = Operate(30) //'~';
	LeftShiftArithmetic  = Operate(31) //'<<';
	RightShiftArithmetic = Operate(32) //'>>';
	RightShiftLogical    = Operate(33) //'>>>';

	/** 其他运算符**/

	//Ternary      = Operate(34) // ? 三元计算
	NullCoalesce = Operate(35) //'??';
	UnaryMinus   = Operate(36) // -/ 负号
	UnaryPlus    = Operate(37) // +/ 正号
	In           = Operate(38) // in
)

type command uint8

const (
	Ellipsis          = command(101)
	Break             = command(102)
	Continue          = command(103)
	Throw             = command(104)
	This              = command(105)
	ExportDefault     = command(106)
	ExportDeclaration = command(107)
	While             = command(108)
	If                = command(109)
	Eos               = command(110)
	Call              = command(111)
	Return            = command(112)
)

// DataType 数据类型
type DataType uint8

const (
	StringType    = DataType(1)  //字符串
	NumberType    = DataType(2)  //数字
	BooleanType   = DataType(3)  //布尔
	NullType      = DataType(4)  //null
	UndefinedType = DataType(5)  //undefined
	BigIntType    = DataType(6)  //bigint
	ObjectType    = DataType(7)  //对象
	ArrayType     = DataType(8)  //数组
	FunctionType  = DataType(9)  //函数
	BufferType    = DataType(10) //字节数组
	ClassType     = DataType(11) //类
)
