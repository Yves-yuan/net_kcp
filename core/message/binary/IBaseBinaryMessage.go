package binary

/*************************************************
  作    者: Yuan
  版    本: v1
  完成日期: 2017-07-11
  功能描述: IBaseBinaryMessage定义将结构写入到buffer的接口
*************************************************/
type IBaseBinaryMessage interface {
	Write(buf *CIOBuffer)
}
