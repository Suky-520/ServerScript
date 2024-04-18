// 这是所有JavaScript计算相关的函数

package ss

import (
	"math"
	"math/big"
)

// + 加
func plus(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(new(big.Int).Add(b1, b2))
		} else if t1 == BigIntType && t2 == NumberType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*Number).Data
			return GoToBigInt(new(big.Int).Add(b1, big.NewInt(int64(b2))))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	//数值 + 数值
	if t1 == NumberType && t2 == NumberType {
		return GoToNumber(JsToNumber(v1) + JsToNumber(v2))
	}
	if t1 == StringType || t2 == StringType {
		return GoToString(JsToString(v1) + JsToString(v2))
	}
	if t1 == ObjectType || t2 == ObjectType || t1 == StringType || t2 == StringType || t1 == ArrayType || t2 == ArrayType || t1 == FunctionType || t2 == FunctionType {
		return GoToString(JsToString(v1) + JsToString(v2))
	}
	if t1 == BooleanType || t2 == BooleanType || t1 == NullType || t2 == NullType || t1 == UndefinedType || t2 == UndefinedType {
		return GoToNumber(JsToNumber(v1) + JsToNumber(v2))
	}

	return GoToString(JsToString(v1) + JsToString(v2))
}

// - 减
func minus(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(big.NewInt(0).Sub(b1, b2))
		} else if t1 == BigIntType && t2 == NumberType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*Number).Data
			return GoToBigInt(new(big.Int).Sub(b1, big.NewInt(int64(b2))))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == ObjectType || t2 == ObjectType || t1 == FunctionType || t2 == FunctionType {
		return NewNaN()
	}
	return GoToNumber(JsToNumber(v1) - JsToNumber(v2))
}

// * 乘
func multiply(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(big.NewInt(0).Mul(b1, b2))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == ObjectType || t2 == ObjectType || t1 == FunctionType || t2 == FunctionType {
		return NewNaN()
	}
	return GoToNumber(JsToNumber(v1) * JsToNumber(v2))
}

// ÷ 除
func divide(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(big.NewInt(0).Div(b1, b2))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == ObjectType || t2 == ObjectType || t1 == FunctionType || t2 == FunctionType {
		return NewNaN()
	}
	return GoToNumber(JsToNumber(v1) / JsToNumber(v2))
}

// % 求余数
func modules(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(big.NewInt(0).Mod(b1, b2))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == ObjectType || t2 == ObjectType || t1 == FunctionType || t2 == FunctionType {
		return NewNaN()
	}
	return GoToNumber(math.Mod(JsToNumber(v1), JsToNumber(v2)))
}

// ** 求指数
func power(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(big.NewInt(0).Exp(b1, b2, nil))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == ObjectType || t2 == ObjectType || t1 == FunctionType || t2 == FunctionType {
		return NewNaN()
	}
	return GoToNumber(math.Pow(JsToNumber(v1), JsToNumber(v2)))
}

// equals == 等于
func equals(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	// NaN 和任何值（包括其本身）都不相等
	if (t1 == NumberType && IsNaN(v1)) || (t2 == NumberType && IsNaN(v2)) {
		return GoToBoolean(false)
	}
	// null 和 undefined 之间是相等的
	if (t1 == NullType && t2 == UndefinedType) || (t1 == UndefinedType && t2 == NullType) {
		return GoToBoolean(true)
	}
	//对象、函数、数组
	if t1 != NumberType && t2 != NumberType && t1 != StringType && t2 != StringType {
		if t1 == ObjectType || t2 == ObjectType || t1 == FunctionType || t2 == FunctionType || t1 == ArrayType || t2 == ArrayType {
			return GoToBoolean(v1 == v2)
		}
	}
	// 布尔值与任何类型比较：布尔值首先转换为数字（true 转为 1，false 转为 0），然后进行比较
	if t1 == BooleanType {
		v1 = GoToNumber(JsToNumber(v1))
		t1 = NumberType
	}
	if t2 == BooleanType {
		v2 = GoToNumber(JsToNumber(v2))
		t2 = NumberType
	}
	if t1 == NumberType && t2 == NumberType {
		return GoToBoolean(JsToNumber(v1) == JsToNumber(v2))
	}
	// 字符串与数字比较：当一个操作数是字符串，另一个是数字时，字符串会尝试转换为数字
	if (t1 == StringType && t2 == NumberType) || (t1 == NumberType && t2 == StringType) {
		return GoToBoolean(JsToNumber(v1) == JsToNumber(v2))
	}
	// 在其他情况下，比较不相等
	return GoToBoolean(JsToString(v1) == JsToString(v2))
}

// === 严格等于
func identityEquals(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	if t1 != t2 {
		return GoToBoolean(false)
	}
	if t1 == ArrayType || t1 == ObjectType || t1 == FunctionType {
		return GoToBoolean(v1 == v2)
	}
	if t1 == NumberType && IsNaN(v1) {
		return GoToBoolean(false)
	}
	if t1 == NumberType && math.Abs(JsToNumber(v1)) == 0 && math.Abs(JsToNumber(v2)) == 0 {
		return GoToBoolean(true)
	}
	return GoToBoolean(JsToString(v1) == JsToString(v2))
}

// != 不等于
func notEquals(v1 JavaScript, v2 JavaScript) JavaScript {
	return GoToBoolean(!equals(v1, v2).(*Boolean).Data)
}

// !== 严格不等于
func identityNotEquals(v1 JavaScript, v2 JavaScript) JavaScript {
	return GoToBoolean(!identityEquals(v1, v2).(*Boolean).Data)
}

// > 大于
func moreThan(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBoolean(b1.Cmp(b2) > 0)
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == StringType && t2 == StringType {
		return GoToBoolean(JsToString(v1) > JsToString(v2))
	}

	v1 = GoToNumber(JsToNumber(v1))
	v2 = GoToNumber(JsToNumber(v2))
	if IsNaN(v1) || IsNaN(v2) {
		return GoToBoolean(false)
	}
	return GoToBoolean(v1.(*Number).Data > v2.(*Number).Data)
}

// < 小于
func lessThan(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBoolean(b1.Cmp(b2) < 0)
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == StringType && t2 == StringType {
		return GoToBoolean(JsToString(v1) < JsToString(v2))
	}
	v1 = GoToNumber(JsToNumber(v1))
	v2 = GoToNumber(JsToNumber(v2))
	if IsNaN(v1) || IsNaN(v2) {
		return GoToBoolean(false)
	}
	return GoToBoolean(v1.(*Number).Data < v2.(*Number).Data)
}

// >= 大于等于
func greaterThanEquals(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			r := b1.Cmp(b2)
			return GoToBoolean(!(r < 0))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}

	if t1 == StringType && t2 == StringType {
		return GoToBoolean(JsToString(v1) >= JsToString(v2))
	}
	v1 = GoToNumber(JsToNumber(v1))
	v2 = GoToNumber(JsToNumber(v2))
	if IsNaN(v1) || IsNaN(v2) {
		return GoToBoolean(false)
	}
	return GoToBoolean(v1.(*Number).Data >= v2.(*Number).Data)
}

// <= 小于等于
func lessThanEquals(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			r := b1.Cmp(b2)
			return GoToBoolean(!(r > 0))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == StringType && t2 == StringType {
		return GoToBoolean(JsToString(v1) <= JsToString(v2))
	}
	v1 = GoToNumber(JsToNumber(v1))
	v2 = GoToNumber(JsToNumber(v2))
	if IsNaN(v1) || IsNaN(v2) {
		return GoToBoolean(false)
	}
	return GoToBoolean(v1.(*Number).Data <= v2.(*Number).Data)
}

// && 逻辑与
func and(v1 JavaScript, v2 JavaScript) JavaScript {

	if !JsToBool(v1) {
		return v1
	}
	return v2
}

// || 逻辑或
func or(v1 JavaScript, v2 JavaScript) JavaScript {
	if JsToBool(v1) {
		return v1
	}
	return v2
}

// ! 逻辑非
func not(v1 JavaScript, v2 JavaScript) JavaScript {
	return GoToBoolean(!JsToBool(v1))
}

// & 按位与
func bitAnd(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(new(big.Int).And(b1, b2))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}

	if t1 == ObjectType || t1 == FunctionType || t2 == ObjectType || t2 == FunctionType {
		return GoToNumber(0)
	}
	return GoToNumber(float64(int64(JsToNumber(v1)) & int64(JsToNumber(v2))))
}

// 按位或：|
func bitOr(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(big.NewInt(0).Or(b1, b2))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == ObjectType || t1 == FunctionType || t2 == ObjectType || t2 == FunctionType {
		return GoToNumber(0)
	}
	return GoToNumber(float64(int64(JsToNumber(v1)) | int64(JsToNumber(v2))))
}

// 按位异或：^
func bitXOr(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(big.NewInt(0).Xor(b1, b2))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == ObjectType || t1 == FunctionType || t2 == ObjectType || t2 == FunctionType {
		return GoToNumber(0)
	}
	return GoToNumber(float64(int64(JsToNumber(v1)) ^ int64(JsToNumber(v2))))
}

// 按位非：~
func bitNot(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType {
		b1 := v1.(*BigInt).Data
		return GoToBigInt(big.NewInt(0).Not(b1))
	}

	if t1 == ObjectType || t1 == FunctionType {
		return GoToNumber(0)
	}
	return GoToNumber(float64(^int64(JsToNumber(v1))))
}

// 左移：<<
func leftShiftArithmetic(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(big.NewInt(0).Lsh(b1, uint(b2.Int64())))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == ObjectType || t1 == FunctionType || t2 == ObjectType || t2 == FunctionType {
		return GoToNumber(0)
	}
	return GoToNumber(float64(int32(JsToNumber(v1)) << int32(JsToNumber(v2))))
}

// 右移：>>
func rightShiftArithmetic(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(big.NewInt(0).Rsh(b1, uint(b2.Int64())))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == ObjectType || t1 == FunctionType || t2 == ObjectType || t2 == FunctionType {
		return GoToNumber(0)
	}
	return GoToNumber(float64(int32(JsToNumber(v1)) >> int32(JsToNumber(v2))))
}

// 无符号右移：>>>
func rightShiftLogical(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	t2 := JsToType(v2).Type
	//校验bigint类型是否和其他类型混合计算
	if t1 == BigIntType || t2 == BigIntType {
		if t1 == BigIntType && t2 == BigIntType {
			b1 := v1.(*BigInt).Data
			b2 := v2.(*BigInt).Data
			return GoToBigInt(big.NewInt(0).Rsh(b1, uint(b2.Int64())))
		} else {
			panic(NewTypeError("Cannot mix BigInt and other types, use explicit conversions"))
		}
	}
	if t1 == ObjectType || t1 == FunctionType || t2 == ObjectType || t2 == FunctionType {
		return GoToNumber(0)
	}
	return GoToNumber(float64(uint32(JsToNumber(v1)) >> (uint(JsToNumber(v2)) % 32)))
}

// - 负号
func unaryMinus(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	if t1 == BigIntType {
		return GoToBigInt(big.NewInt(0).Neg(v1.(*BigInt).Data))
	}
	return GoToNumber(-JsToNumber(v1))
}

// + 正号
func unaryPlus(v1 JavaScript, v2 JavaScript) JavaScript {
	t1 := JsToType(v1).Type
	if t1 == BigIntType {
		return GoToBigInt(big.NewInt(0).Abs(v1.(*BigInt).Data))
	}
	if t1 == NumberType {
		if JsToNumber(v1) < 0 {
			return GoToNumber(-JsToNumber(v1))
		}
	}
	return GoToNumber(JsToNumber(v1))
}

// in 运算符：检测对象是否具有给定的属性。
func in(v1 JavaScript, v2 JavaScript) JavaScript {
	if _, ok := JsGetProperty(v2, JsToString(v1)).(*Undefined); ok {
		return GoToBoolean(false)
	}
	return GoToBoolean(true)
}

// DefaultCalculate 默认运算
func DefaultCalculate(v1 JavaScript, v2 JavaScript, cal Operate) JavaScript {
	switch cal {
	case Plus:
		return plus(v1, v2)
	case Minus:
		return minus(v1, v2)
	case Multiply:
		return multiply(v1, v2)
	case Divide:
		return divide(v1, v2)
	case Modulus:
		return modules(v1, v2)
	case Power:
		return power(v1, v2)
	case Equals:
		return equals(v1, v2)
	case IdentityEquals:
		return identityEquals(v1, v2)
	case NotEquals:
		return notEquals(v1, v2)
	case IdentityNotEquals:
		return identityNotEquals(v1, v2)
	case MoreThan:
		return moreThan(v1, v2)
	case LessThan:
		return lessThan(v1, v2)
	case GreaterThanEquals:
		return greaterThanEquals(v1, v2)
	case LessThanEquals:
		return lessThanEquals(v1, v2)
	case And:
		return and(v1, v2)
	case Or:
		return or(v1, v2)
	case Not:
		return not(v1, v2)
	case BitAnd:
		return bitAnd(v1, v2)
	case BitOr:
		return bitOr(v1, v2)
	case BitXOr:
		return bitXOr(v1, v2)
	case BitNot:
		return bitNot(v1, v2)
	case LeftShiftArithmetic:
		return leftShiftArithmetic(v1, v2)
	case RightShiftArithmetic:
		return rightShiftArithmetic(v1, v2)
	case RightShiftLogical:
		return rightShiftLogical(v1, v2)
	case UnaryPlus:
		return unaryPlus(v1, v2)
	case UnaryMinus:
		return unaryMinus(v1, v2)
	case In:
		return in(v1, v2)
	default:
		panic(NewTypeError("运算符不支持"))
	}
}
