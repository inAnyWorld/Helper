package gostring

import (
	"bytes"
	"crypto/md5"
	_ "crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//ByteToString byte转string
func ByteToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	str := s.String()
	return &str
}

// ConvString 转换为string
func ConvString(i interface{}) (s string) {
	switch i.(type) {
	case int:
		s = strconv.Itoa(i.(int))
	case uint64:
		s = strconv.FormatUint(i.(uint64), 10)
	case int64:
		s = strconv.FormatInt(i.(int64), 10)
	case float64:
		s = strconv.FormatFloat(i.(float64), 'f', -1, 64)
	case string:
		s = i.(string)
	default:
		s = ""
	}
	return
}

// Base64Encode base64编码
func Base64Encode(encodeStr string ) string {
	base64Str := encodeStr
	base64Byte := []byte(base64Str)
	return base64.StdEncoding.EncodeToString(base64Byte)
}

// StructName 获取结构体名
func StructName(i interface{}) string {
	structName := reflect.TypeOf(i).Elem().Name()
	return structName
}

// CalculatePercentage 计算百分比
func CalculatePercentage(compare int, toBeCompare int) string {
	if toBeCompare == 0 {
		return "0.00%"
	}
	if compare == toBeCompare {
		return "0.00%"
	}
	percentage := float64(compare - toBeCompare) / float64(toBeCompare)
	percentageValue, percentageErr := strconv.ParseFloat(fmt.Sprintf("%.4f", percentage), 64)
	if percentageErr != nil {
		return "0.00%"
	}
	return fmt.Sprintf("%v%%", percentageValue * 100)
}

// IncomePercentage 计算收入百分比(暂时不会使用此字段)
func IncomePercentage() string {
	return "0.00%"
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	letterRunes := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RandomString2 1字母,2数值,3字母+数值,4字母+数值+特殊字符,5仅中文
func RandomString2(length uint8, sType int) string {
	if length == 0 {
		return ""
	}
	var letter []rune
	alphas := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specials := "~!@#$%^&*()_+{}:|<>?`-=;,."

	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Nanosecond)

	switch sType {
	case 1:
		letter = []rune(alphas)
	case 2:
		letter = []rune(numbers)
	case 3:
		letter = []rune(alphas + numbers)
	case 4:
		letter = []rune(alphas + numbers + specials)
	case 5:
		chineses := "们以我到他会作时要动国产的一是工就年阶义发成部民可出能方进在了不和有大这主中人上为来分生对于学下级地个用同行面说种过命度革而多子后自社加小机也经力线本电高量长党得实家定深法表着水理化争现所二起政三好十战无农使性前等反体合斗路图把结第里正新开论之物从当两些还天资事队批点育重其思与间内去因件日利相由压员气业代全组数果期导平各基或月毛然如应形想制心样干都向变关问比展那它最及外没看治提五解系林者米群头意只明四道马认次文通但条较克又公孔领军流入接席位情运器并飞原油放立题质指建区验活众很教决特此常石强极土少已根共直团统式转别造切九你取西持总料连任志观调七么山程百报更见必真保热委手改管处己将修支识病象几先老光专什六型具示复安带每东增则完风回南广劳轮科北打积车计给节做务被整联步类集号列温装即毫知轴研单色坚据速防史拉世设达尔场织历花受求传口断况采精金界品判参层止边清至万确究书术状厂须离再目海交权且儿青才证低越际八试规斯近注办布门铁需走议县兵固除般引齿千胜细影济白格效置推空配刀叶率述今选养德话查差半敌始片施响收华觉备名红续均药标记难存测士身紧液派准斤角降维板许破述技消底床田势端感往神便贺村构照容非搞亚磨族火段算适讲按值美态黄易彪服早班麦削信排台声该击素张密害侯草何树肥继右属市严径螺检左页抗苏显苦英快称坏移约巴材省黑武培著河帝仅针怎植京助升王眼她抓含苗副杂普谈围食射源例致酸旧却充足短划剂宣环落首尺波承粉践府鱼随考刻靠够满夫失包住促枝局菌杆周护岩师举曲春元超负砂封换太模贫减阳扬江析亩木言球朝医校古呢稻宋听唯输滑站另卫字鼓刚写刘微略范供阿块某功套友限项余倒卷创律雨让骨远帮初皮播优占死毒圈伟季训控激找叫云互跟裂粮粒母练塞钢顶策双留误础吸阻故寸盾晚丝女散焊功株亲院冷彻弹错散商视艺灭版烈零室轻血倍缺厘泵察绝富城冲喷壤简否柱李望盘磁雄似困巩益洲脱投送奴侧润盖挥距触星松送获兴独官混纪依未突架宽冬章湿偏纹吃执阀矿寨责熟稳夺硬价努翻奇甲预职评读背协损棉侵灰虽矛厚罗泥辟告卵箱掌氧恩爱停曾溶营终纲孟钱待尽俄缩沙退陈讨奋械载胞幼哪剥迫旋征槽倒握担仍呀鲜吧卡粗介钻逐弱脚怕盐末阴丰雾冠丙街莱贝辐肠付吉渗瑞惊顿挤秒悬姆烂森糖圣凹陶词迟蚕亿矩康遵牧遭幅园腔订香肉弟屋敏恢忘编印蜂急拿扩伤飞露核缘游振操央伍域甚迅辉异序免纸夜乡久隶缸夹念兰映沟乙吗儒杀汽磷艰晶插埃燃欢铁补咱芽永瓦倾阵碳演威附牙芽永瓦斜灌欧献顺猪洋腐请透司危括脉宜笑若尾束壮暴企菜穗楚汉愈绿拖牛份染既秋遍锻玉夏疗尖殖井费州访吹荣铜沿替滚客召旱悟刺脑措贯藏敢令隙炉壳硫煤迎铸粘探临薄旬善福纵择礼愿伏残雷延烟句纯渐耕跑泽慢栽鲁赤繁境潮横掉锥希池败船假亮谓托伙哲怀割摆贡呈劲财仪沉炼麻罪祖息车穿货销齐鼠抽画饲龙库守筑房歌寒喜哥洗蚀废纳腹乎录镜妇恶脂庄擦险赞钟摇典柄辩竹谷卖乱虚桥奥伯赶垂途额壁网截野遗静谋弄挂课镇妄盛耐援扎虑键归符庆聚绕摩忙舞遇索顾胶羊湖钉仁音迹碎伸灯避泛亡答勇频皇柳哈揭甘诺概宪浓岛袭谁洪谢炮浇斑讯懂灵蛋闭孩释乳巨徒私银伊景坦累匀霉杜乐勒隔弯绩招绍胡呼痛峰零柴簧午跳居尚丁秦稍追梁折耗碱殊岗挖氏刃剧堆赫荷胸衡勤膜篇登驻案刊秧缓凸役剪川雪链渔啦脸户洛孢勃盟买杨宗焦赛旗滤硅炭股坐蒸凝竟陷枪黎救冒暗洞犯筒您宋弧爆谬涂味津臂障褐陆啊健尊豆拔莫抵桑坡缝警挑污冰柬嘴啥饭塑寄赵喊垫丹渡耳刨虎笔稀昆浪萨茶滴浅拥穴覆伦娘吨浸袖珠雌妈紫戏塔锤震岁貌洁剖牢锋疑霸闪埔猛诉刷狠忽灾闹乔唐漏闻沈熔氯荒茎男凡抢像浆旁玻亦忠唱蒙予纷捕锁尤乘乌智淡允叛畜俘摸锈扫毕璃宝芯爷鉴秘净蒋钙肩腾枯抛轨堂拌爸循诱祝励肯酒绳穷塘燥泡袋朗喂铝软渠颗惯贸粪综墙趋彼届墨碍启逆卸航衣孙龄岭骗休借"
		letter = []rune(chineses)
	default:
		letter = []rune(alphas)
	}

	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

// GetBetweenStr 截取指定字符串
func GetBetweenStr(str, start string, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

// Substr 字符串截取
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// SubStrLeftOrRight 从左右截取指定字符串
func SubStrLeftOrRight(str string, target string, direction string, hasPos bool) (string, error) {
	pos := strings.Index(str, target)
	if pos == -1 {
		return "", errors.New("pos -1")
	}
	if direction == "left" {
		if hasPos == true {
			pos += 1
		}
		return str[:pos], nil
	}
	if direction == "right" {
		if hasPos == false {
			pos += 1
		}
		return str[pos:], nil
	}
	return "", errors.New("parameter error")
}

// MD5 md5字符串
func MD5(md5Str string) string {
	md5Byte := []byte(md5Str)
	md5Sum := md5.Sum(md5Byte)
	return strings.ToUpper(fmt.Sprintf("%x", md5Sum))
}

// UcFirst 首字符大写
func UcFirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}

// LcFirst 首字母小写
func LcFirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

// UcWords 将字符串中每个词的首字母转换为大写
func UcWords(str string) string {
	return strings.Title(str)
}

// LcWords 单词首字母小写
func LcWords(str string) string {
	buf := &bytes.Buffer{}
	lastIsSpace := true
	for _, r := range str {
		if unicode.IsLetter(r) {
			if lastIsSpace {
				r = unicode.ToLower(r)
			}

			lastIsSpace = false
		} else {
			lastIsSpace = false
			if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsMark(r) {
				lastIsSpace = true
			}
		}

		buf.WriteRune(r)
	}

	return buf.String()
}

// Shuffle 打乱字符串
func Shuffle(str string) string {
	if str == "" {
		return str
	}

	runes := []rune(str)
	index := 0

	for i := len(runes) - 1; i > 0; i-- {
		index = rand.Intn(i + 1)

		if i != index {
			runes[i], runes[index] = runes[index], runes[i]
		}
	}

	return string(runes)
}

// Strops2 查找字符串首次出现的位置,找不到时返回-1.haystack在该字符串中进行查找,needle要查找的字符串,offset起始位置.
func Strops2(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// Strips2  查找字符串首次出现的位置（不区分大小写）,找不到时返回-1. haystack在该字符串中进行查找,needle要查找的字符串,offset起始位置
func Strips2(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	haystack = haystack[offset:]
	pos := strings.Index(strings.ToLower(haystack), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// Strops 查找指定字符串在目标字符串中最后一次出现的位置
func Strops(haystack, needle string, offset int) int {
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(haystack, needle)
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// Strips 查找指定字符串在目标字符串中最后一次出现的位置（不区分大小写
func Strips(haystack, needle string, offset int) int {
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(strings.ToLower(haystack), strings.ToLower(needle))
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// Md5Str 计算字符串的 MD5 散列值.
func Md5Str(str []byte, length uint8) []byte {
	var res []byte
	h := md5.New()
	h.Write(str)

	hBytes := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(dst, hBytes)
	if length > 0 && length < 32 {
		res = dst[:length]
	} else {
		res = dst
	}

	return res
}