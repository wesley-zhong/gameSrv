Public buffWb As Workbook
Public buffEffectWb As Workbook
Public buffCondWb As Workbook

Public buffSheet As Worksheet
Public buffEffectSheet As Worksheet
Public buffCondSheet As Worksheet

' 属性Cell
Public PropertyTypeCell As Range

' 特殊效果Cell
Public SpecialEffectType As Range

' 特殊条件Cell
Public SpecialConditionCell As Range

' 触发需要的MTG_ID
Public TriggerMontage As Range
'buff Trigger Events
Public StartOfTriggerEvents As Range

Public BuffDeathEventCell As Range
Public RemoveTriggerEvents As Range

Public BuffLength As Integer
Public adwBuffEffect As Range
Public condIds As Range

' 新增：EffectTarget列
Public EffectTargetCell As Range

' 新增：BuffDeathCounter列
Public BuffDeathCounterCell As Range

Private Sub ComandButton_Click()
    'Just Test
    'cellValue = Cells(1, 1).Value
    'Range("A2").Value = cellValue
    
    '打开Buff+BuffEffect
    Set buffEffectWb = Workbooks.Open(ThisWorkbook.Path & "\buffeffect_Template.xlsx")
    Set buffWb = Workbooks.Open(ThisWorkbook.Path & "\buff_Template.xlsx")
    Set buffCondWb = Workbooks.Open(ThisWorkbook.Path & "\buffcond_Template.xlsx")

    
    Set buffSheet = buffWb.Worksheets("buff")
    Set buffEffectSheet = buffEffectWb.Worksheets("buffeffect")
    Set buffCondSheet = buffCondWb.Worksheets("condition")
    
        
    '将目标工作表的所有单元格格式设置为文本
    buffSheet.Cells.NumberFormat = "@"
    buffEffectSheet.Cells.NumberFormat = "@"
    buffCondSheet.Cells.NumberFormat = "@"

    
    ' 找到特殊的buffEffect
    Set PropertyTypeCell = Cells.Find("PropertyType", LookAt:=xlWhole)
    Set SpecialEffectType = Cells.Find("SpecialEffectType", LookAt:=xlWhole)
    Set TriggerMontage = Cells.Find("TriggerMontage", LookAt:=xlWhole)
    Set SpecialConditionCell = TriggerMontage.Offset(0, 5)

    Set BuffDeathEventCell = Cells.Find("BuffDeathEvent", LookAt:=xlWhole)

    ' 新增：查找EffectTarget列
    Set EffectTargetCell = Cells.Find("EffectTarget", LookAt:=xlWhole)

    ' 新增：查找BuffDeathCounter列
    Set BuffDeathCounterCell = Cells.Find("BuffDeathCounter", LookAt:=xlWhole)

    '映射
    Dim tempRow As Integer
    tempRow = 6

    ' 找adwBuffEffect
    Set adwBuffEffect = buffSheet.Cells.Find("effectIds", LookAt:=xlWhole)
    Set condIds = buffSheet.Cells.Find("condIds", LookAt:=xlWhole)
    Set StartOfTriggerEvents = buffSheet.Cells.Find("StartOfTrigger", LookAt:=xlWhole).Offset(4 - tempRow)
    Set RemoveTriggerEvents = buffSheet.Cells.Find("RemoveTrigger", LookAt:=xlWhole).Offset(4 - tempRow)
    
    ' 以防修改中的意外情况
    BuffLength = 0

    For i = 1 To 22
        Dim nameValue
        Set nameValue = Cells(3, i)

        Dim tempCol As Integer
        tempCol = i

        'buff 部分映射
        temp = FindAndReplace_buff(nameValue, tempRow, tempCol)

    Next i

    For i = 41 To 80
        Dim nameValue1
        Set nameValue1 = Cells(3, i)

        Dim tempCol1 As Integer
        tempCol1 = i

        'buff 部分映射
        temp = FindAndReplace_buff(nameValue1, tempRow, tempCol1)

    Next i


    'buff Effect部分映射
    temp1 = Create_buffEffect(tempRow)

    '插入Trigger
    temp2 = Insert_Trigger(tempRow)
    temp3 = Insert_RemoveTrigger(tempRow)

    ' 处理第一组条件
    Dim lastEffectIndex As Long
    lastEffectIndex = Create_buffEffectCondition(tempRow)
    ' 处理第二组条件（使用第一组的lastEffectIndex作为起始）
    lastEffectIndex = Create_buffEffectCondition2(tempRow, lastEffectIndex)

    temp4 = CombineMultiEffectAndCondition(tempRow)

    Application.DisplayAlerts = False
    '保存
    buffWb.SaveAs ThisWorkbook.Path & "\buff.xlsx"
    buffEffectWb.SaveAs ThisWorkbook.Path & "\buffEffect.xlsx"
    buffCondWb.SaveAs ThisWorkbook.Path & "\buffCond.xlsx"
    '关闭
    buffEffectWb.Close
    buffWb.Close
    buffCondWb.Close

    'Open Self
    ThisWorkbook.Save
    'test = Workbooks.Open(ThisWorkbook.Path & "\buff_design.xlsm")
    Application.DisplayAlerts = True
End Sub

Function CombineMultiEffectAndCondition(tempRow As Integer)
    Dim IsCombining As Boolean
    Dim FinalRow As Long
    Dim rowValue As Range
    Dim condIdsValue As String
    Dim effectValue As String
    Dim i As Long

    IsCombining = False
    FinalRow = 0 '合并到哪一行
    For i = tempRow To BuffLength
        Set rowValue = Cells(i, 1)
        '特殊标记
        If (rowValue = "###") Then
            If (IsCombining) Then
                ' 不做任何操作，继续合并
            Else
                IsCombining = True
                FinalRow = i - 1 ' 合并到第一次出现的上一行
            End If
            condIdsValue = condIds.Offset(i - tempRow).Value
            If (condIdsValue <> "" And condIdsValue <> "0") Then
                If condIds.Offset(FinalRow - tempRow).Value = "" Or condIds.Offset(FinalRow - tempRow).Value = "0" Then
                    condIds.Offset(FinalRow - tempRow).Value = condIdsValue
                Else
                    condIds.Offset(FinalRow - tempRow).Value = condIds.Offset(FinalRow - tempRow).Value & "," & condIdsValue
                End If
            End If
            effectValue = adwBuffEffect.Offset(i - tempRow).Value
            If (effectValue <> "") Then
                If adwBuffEffect.Offset(FinalRow - tempRow).Value = "" Then
                    adwBuffEffect.Offset(FinalRow - tempRow).Value = effectValue
                Else
                    adwBuffEffect.Offset(FinalRow - tempRow).Value = adwBuffEffect.Offset(FinalRow - tempRow).Value & "," & effectValue
                End If
            End If
        Else
            IsCombining = False
        End If
    Next i

    ' 删除特殊标记的行
    For i = BuffLength To tempRow Step -1 '倒序循环
        Set rowValue = Cells(i, 1)
        '特殊标记
        If (rowValue = "###") Or (rowValue = "$$$") Then
            buffSheet.Rows(i - tempRow + 4).Delete
        End If
    Next
End Function

Function Create_buffEffectCondition(tempRow As Integer) As Long
    Dim EffectIndex As Long
    EffectIndex = 5  ' 初始Index
    Const OffsetNumber As Long = 100000

    For i = tempRow To BuffLength
        EffectStr = ""
        Set rowValue = SpecialConditionCell.Offset(i - 1, 0)
        
        If (rowValue = "") Or (rowValue = "0") Or (rowValue = "NONE") Then
            If i <= BuffLength - 2 Then
                If condIds.Offset(i - tempRow + 2).Value = "" Then
                    condIds.Offset(i - tempRow + 2).Value = "0"
                End If
            End If
            GoTo ContinueLoop
        End If
        
        Set IDCell = buffCondSheet.Cells(EffectIndex + 1, 2)
        IDCell.Value = (OffsetNumber + EffectIndex)

        IDCell.Offset(0, 1).Value = 0

        IDCell.Offset(0, 2).Value = rowValue
        IDCell.Offset(0, 3).Value = rowValue.Offset(0, 1)
        IDCell.Offset(0, 4).Value = rowValue.Offset(0, 2)

        ' 直接设置值（第二组会追加）
        condIds.Offset(i - tempRow + 2).Value = (OffsetNumber + EffectIndex)
        EffectIndex = EffectIndex + 1

ContinueLoop:
    Next i

    ' 返回最后使用的索引
    Create_buffEffectCondition = EffectIndex
End Function

Function Create_buffEffectCondition2(tempRow As Integer, startIndex As Long) As Long
    Dim EffectIndex As Long
    EffectIndex = startIndex
    Const OffsetNumber As Long = 200000

    Dim SecondSpecialConditionCell As Range
    Set SecondSpecialConditionCell = SpecialConditionCell.Offset(0, 4)
    
    For i = tempRow To BuffLength
        EffectStr = ""
        Set rowValue = SecondSpecialConditionCell.Offset(i - 1, 0)
        
        If (rowValue = "") Or (rowValue = "0") Or (rowValue = "NONE") Then
            GoTo ContinueLoop
        End If
        
        Set IDCell = buffCondSheet.Cells(EffectIndex + 1, 2)
        IDCell.Value = (OffsetNumber + EffectIndex)
        IDCell.Offset(0, 1).Value = 1

        IDCell.Offset(0, 2).Value = rowValue
        IDCell.Offset(0, 3).Value = rowValue.Offset(0, 1)
        IDCell.Offset(0, 4).Value = rowValue.Offset(0, 2)

        Dim existingValue As String
        existingValue = condIds.Offset(i - tempRow + 2).Value

        If existingValue = "" Or existingValue = "0" Then
            condIds.Offset(i - tempRow + 2).Value = (OffsetNumber + EffectIndex)
        Else
            condIds.Offset(i - tempRow + 2).Value = existingValue & "," & (OffsetNumber + EffectIndex)
        End If

        EffectIndex = EffectIndex + 1

ContinueLoop:
    Next i

    Create_buffEffectCondition2 = EffectIndex
End Function

Function Insert_RemoveTrigger(tempRow As Integer)
    AdditiveTriggerNum = -1 'trigger会增加行数
    For i = tempRow To BuffLength
        num = 0
        Set rowValue = BuffDeathEventCell.Offset(i - 1, 0)
        If rowValue = "" Then
            GoTo ContinueLoop2
        End If
        If num > 0 Then
            buffSheet.Rows(RemoveTriggerEvents.Offset(i - 2 + AdditiveTriggerNum, 0).Row).Insert Shift:=xlDown
        End If

        Set IDCell = RemoveTriggerEvents.Offset(i - 2 + AdditiveTriggerNum, 0)
        IDCell.Offset(0, 1).Value = rowValue
        If rowValue = "挂载时" Then
            IDCell.Value = "CommonTriggerEvent"
        ElseIf rowValue = "进入编队" Then
            IDCell.Value = "TeamTriggerEvent"
        ElseIf rowValue = "离开编队" Then
            IDCell.Value = "TeamTriggerEvent"
        ElseIf rowValue = "命中时" Then
            IDCell.Value = "HitedTriggerEvent"
        ElseIf rowValue = "造成伤害时" Then
            IDCell.Value = "DoDamageTriggerEvent"
        ElseIf rowValue = "暴击时" Then
            IDCell.Value = "DoCriticalDamageTriggerEvent"
        ElseIf rowValue = "受击时" Then
            IDCell.Value = "BeHitedTriggerEvent"
        ElseIf rowValue = "受到伤害时" Then
            IDCell.Value = "BeDamageTriggerEvent"
        ElseIf rowValue = "角色属性达到指定值" Then
            IDCell.Value = "PropChangeTriggerEvent"
        ElseIf rowValue = "状态改变时" Then
            IDCell.Value = "StatusChangeTriggerEvent"
        ElseIf rowValue = "释放蒙太奇ByID" Then
            IDCell.Value = "MtgIdTriggerEvent"
        ElseIf rowValue = "释放蒙太奇ByAbilityTag" Then
            IDCell.Value = "MtgTagTriggerEvent"
        ElseIf rowValue = "永不触发" Then
            IDCell.Value = "NeverTriggerEvent"
        ElseIf rowValue = "命中时ByAttackDateID" Then
            IDCell.Value = "HitedTriggerEvent"
        ElseIf rowValue = "角色属性达到指定比例" Then
            IDCell.Value = "PropChangeTriggerEvent"
        ElseIf rowValue = "命中时BySkillType" Then
            IDCell.Value = "HitedSkillTypeTriggerEvent"
        ElseIf rowValue = "完美闪避时" Then
            IDCell.Value = "DefaultTrueTrigger"
        ElseIf rowValue = "拼刀时" Then
            IDCell.Value = "DefaultTrueTrigger"
        ElseIf rowValue = "击杀怪物时" Then
            IDCell.Value = "KilledMonsterEvent"
        ElseIf rowValue = "主角周围有怪物进入虚弱" Then
            IDCell.Value = "AroundHasMonsterBeWeak"
        ElseIf rowValue = "蒙太奇结束时" Then
            IDCell.Value = "MontageEndTriggerEvent"
        ElseIf rowValue = "处决结束时" Then
            IDCell.Value = "ExecutionEnd"
        ElseIf rowValue = "施加结晶时ByAttackID" Then
            IDCell.Value = "DoAbnormalClosing"
        ElseIf rowValue = "造成治疗时" Then
            IDCell.Value = "DoDamageTriggerEvent"
        ElseIf rowValue = "受到治疗时" Then
            IDCell.Value = "BeDamageTriggerEvent"
        End If
        IDCell.Offset(0, 4).Value = rowValue.Offset(0, 1)

        If IDCell.Offset(0, 2).Value = "" Then
            IDCell.Offset(0, 2).Value = 0
        End If
        If IDCell.Offset(0, 3).Value = "" Then
            IDCell.Offset(0, 3).Value = 0
        End If
        If IDCell.Offset(0, 4).Value = "" Then
            IDCell.Offset(0, 4).Value = 0
        End If
        IDCell.Offset(0, 5).Value = rowValue.Offset(0, 2)
        IDCell.Offset(0, 6).Value = rowValue.Offset(0, 3)
        num = num + 1

ContinueLoop2:
    Next i

    ' ========== 新增：复制 BuffDeathCounter 列到 buffSheet ==========
    Dim bcCounterCol As Range
    Dim bcTargetRow As Long
    Dim bcValue As Variant
    Dim j As Long

    ' 查找 buffSheet 中的 "BuffDeathCounter" 列
    Set bcCounterCol = buffSheet.Cells.Find("BuffDeathCounter", LookAt:=xlWhole)
    If Not bcCounterCol Is Nothing Then
        ' 遍历所有数据行
        For j = tempRow To BuffLength
            bcValue = BuffDeathCounterCell.Offset(j - 1, 0).Value
            ' 严格按照原有偏移逻辑计算目标行号（与 RemoveTrigger 数据行完全对齐）
            bcTargetRow = RemoveTriggerEvents.Offset(j - 3).Row
            buffSheet.Cells(bcTargetRow, bcCounterCol.Column).Value = bcValue
        Next j
    End If
    ' ================================================================
End Function

Function Insert_Trigger(tempRow As Integer)
    AdditiveTriggerNum = -1 'trigger会增加行数
    For i = tempRow To BuffLength
        num = 0
        '触发条件类型部分
        Set rowValue = TriggerMontage.Offset(i - 1, 1)
        If rowValue = "" Then
            GoTo ContinueLoop2
        End If
        If num > 0 Then
            buffSheet.Rows(StartOfTriggerEvents.Offset(i - 2 + AdditiveTriggerNum, 0).Row).Insert Shift:=xlDown
        End If

        Set IDCell = StartOfTriggerEvents.Offset(i - 2 + AdditiveTriggerNum, 0)
        IDCell.Offset(0, 1).Value = rowValue
        If rowValue = "挂载时" Then
            IDCell.Value = "CommonTriggerEvent"
        ElseIf rowValue = "进入编队" Then
            IDCell.Value = "TeamTriggerEvent"
        ElseIf rowValue = "离开编队" Then
            IDCell.Value = "TeamTriggerEvent"
        ElseIf rowValue = "命中时" Then
            IDCell.Value = "HitedTriggerEvent"
        ElseIf rowValue = "造成伤害时" Then
            IDCell.Value = "DoDamageTriggerEvent"
        ElseIf rowValue = "暴击时" Then
            IDCell.Value = "DoCriticalDamageTriggerEvent"
        ElseIf rowValue = "受击时" Then
            IDCell.Value = "BeHitedTriggerEvent"
        ElseIf rowValue = "受到伤害时" Then
            IDCell.Value = "BeDamageTriggerEvent"
        ElseIf rowValue = "角色属性达到指定值" Then
            IDCell.Value = "PropChangeTriggerEvent"
        ElseIf rowValue = "状态改变时" Then
            IDCell.Value = "StatusChangeTriggerEvent"
        ElseIf rowValue = "释放蒙太奇ByID" Then
            IDCell.Value = "MtgIdTriggerEvent"
        ElseIf rowValue = "释放蒙太奇ByAbilityTag" Then
            IDCell.Value = "MtgTagTriggerEvent"
        ElseIf rowValue = "永不触发" Then
            IDCell.Value = "NeverTriggerEvent"
        ElseIf rowValue = "命中时ByAttackDateID" Then
            IDCell.Value = "HitedTriggerEvent"
        ElseIf rowValue = "角色属性达到指定比例" Then
            IDCell.Value = "PropChangeTriggerEvent"
        ElseIf rowValue = "命中时BySkillType" Then
            IDCell.Value = "HitedSkillTypeTriggerEvent"
        ElseIf rowValue = "完美闪避时" Then
            IDCell.Value = "DefaultTrueTrigger"
        ElseIf rowValue = "拼刀时" Then
            IDCell.Value = "DefaultTrueTrigger"
        ElseIf rowValue = "击杀怪物时" Then
            IDCell.Value = "KilledMonsterEvent"
        ElseIf rowValue = "主角周围有怪物进入虚弱" Then
            IDCell.Value = "AroundHasMonsterBeWeak"
        ElseIf rowValue = "蒙太奇结束时" Then
            IDCell.Value = "MontageEndTriggerEvent"
        ElseIf rowValue = "处决结束时" Then
            IDCell.Value = "ExecutionEnd"
        ElseIf rowValue = "施加结晶时ByAttackID" Then
            IDCell.Value = "DoAbnormalClosing"
        ElseIf rowValue = "造成治疗时" Then
            IDCell.Value = "DoDamageTriggerEvent"
        ElseIf rowValue = "受到治疗时" Then
            IDCell.Value = "BeDamageTriggerEvent"
        End If

        IDCell.Offset(0, 4).Value = rowValue.Offset(0, 1)
        If IDCell.Offset(0, 4).Value = "" Then
            IDCell.Offset(0, 4).Value = 0
        End If
        IDCell.Offset(0, 5).Value = rowValue.Offset(0, 2)
        IDCell.Offset(0, 6).Value = rowValue.Offset(0, 3)
        num = num + 1

ContinueLoop2:
    Next i
End Function

Function FindAndReplace_buff(key, tempRow As Integer, tempCol As Integer)
    Set thatOne = buffSheet.Cells.Find(key, LookAt:=xlWhole)
    If thatOne Is Nothing Then
        Exit Function
    End If
    FindAndReplace_buff = thatOne
    For i = tempRow To 10000
        rowValue = Cells(i + 2, tempCol)
        tempValue = rowValue
        If rowValue = "" Then
            If i >= BuffLength - 1 Then
                BuffLength = i + 1
                Exit For
            End If
            tempValue = Cells(tempRow + 1, tempCol)
        End If
        thatOne.Offset(i - tempRow + 5, 0).Value = tempValue
    Next i
End Function

' 新增：将EffectTarget中文转换为枚举名称
Function ConvertEffectTarget(target As Variant) As String
    If target = "" Then
        ConvertEffectTarget = ""
        Exit Function
    End If

    Select Case target
        Case "自己"
            ConvertEffectTarget = "MYSELF"
        Case "交互对象"
            ConvertEffectTarget = "ATTACK_TARGET"
        Case "编队队友"
            ConvertEffectTarget = "TEAM_EXCLUDE_ME"
        Case "编队包含自己"
            ConvertEffectTarget = "TEAM_INCLUDE_ME"
        Case "周围队友"
            ConvertEffectTarget = "AROUND_EXCLUDE_ME"
        Case "周围包含自己"
            ConvertEffectTarget = "AROUND_INCLUDE_ME"
        Case "后台队友"
            ConvertEffectTarget = "BACKSTAGE_EXCLUDE_ME"
        Case "后台包含自己"
            ConvertEffectTarget = "BACKSTAGE_INCLUDE_ME"
        Case "站场队友"
            ConvertEffectTarget = "ONSTAGE_EXCLUDE_ME"
        Case "站场包含自己"
            ConvertEffectTarget = "ONSTAGE_INCLUDE_ME"
        Case Else
            ConvertEffectTarget = ""
    End Select
End Function

Function Create_buffEffect(tempRow As Integer)
    EffectIndex = tempRow - tempRow + 5
    OffsetNumber = 100000

    Dim currentCnfId As Variant
    Dim cachedTarget As Variant

    For i = tempRow To BuffLength
        EffectStr = ""

        currentCnfId = Cells(i, 1).Value

        Dim rawTarget As Variant
        If Not EffectTargetCell Is Nothing Then
            rawTarget = EffectTargetCell.Offset(i - 1, 0).Value
        Else
            rawTarget = ""
        End If

        Dim finalTarget As Variant
        If rawTarget <> "" Then
            finalTarget = rawTarget
            cachedTarget = finalTarget
        Else
            finalTarget = cachedTarget
        End If

        ' 将最终的中文目标转换为枚举数字字符串
        Dim targetEnum As String
        targetEnum = ConvertEffectTarget(finalTarget)

        ' 处理普通属性效果
        Set rowValue = PropertyTypeCell.Offset(i - 1, 0)
        If Not ((rowValue = "") Or (rowValue = "0") Or (rowValue = "NONE")) Then
            Set IDCell = buffEffectSheet.Cells(EffectIndex, 2)
            IDCell.Value = (OffsetNumber + EffectIndex)
            IDCell.Offset(0, 1).Value = targetEnum   ' 写入枚举值
            IDCell.Offset(0, 2).Value = "ModAvatarProps"
            IDCell.Offset(0, 3).Value = rowValue.Offset(0, 0)
            Select Case rowValue.Offset(0, 2).Value
                Case "固定值"
                    IDCell.Offset(0, 4).Value = "加"
                Case "万分比"
                    IDCell.Offset(0, 4).Value = "增加(万分比)"
                Case Else
                    IDCell.Offset(0, 4).Value = "加"
            End Select
            IDCell.Offset(0, 5).Value = rowValue.Offset(0, 1)
            If EffectStr = "" Then
                EffectStr = (OffsetNumber + EffectIndex)
            Else
                EffectStr = EffectStr & "," & (OffsetNumber + EffectIndex)
            End If
            adwBuffEffect.Offset(i - tempRow + 2).Value = EffectStr
            EffectIndex = EffectIndex + 1
        End If

        ' 处理特殊效果
        Set rowValue = SpecialEffectType.Offset(i - 1, 0)
        If Not ((rowValue = "") Or (rowValue = "0") Or (rowValue = "NONE")) Then
            Set IDCell = buffEffectSheet.Cells(EffectIndex, 2)
            IDCell.Value = (OffsetNumber + EffectIndex)
            IDCell.Offset(0, 1).Value = targetEnum   ' 写入枚举值

            If rowValue = "ModAvatarPropsForEver" Then
                IDCell.Offset(0, 2).Value = rowValue              ' D列: BuffEffectOpt
                IDCell.Offset(0, 3).Value = rowValue.Offset(0, 1) ' E列: 属性名
                Select Case rowValue.Offset(0, 3).Value
                    Case "固定值"
                        IDCell.Offset(0, 4).Value = "加"          ' F列: 操作符
                    Case "万分比"
                        IDCell.Offset(0, 4).Value = "增加(万分比)"
                    Case Else
                        IDCell.Offset(0, 4).Value = "加"
                End Select
                IDCell.Offset(0, 5).Value = rowValue.Offset(0, 2) ' G列: 数值

            ElseIf rowValue = "ModAvatarProps" Then
                IDCell.Offset(0, 2).Value = rowValue
                IDCell.Offset(0, 3).Value = rowValue.Offset(0, 1)
                Select Case rowValue.Offset(0, 3).Value
                    Case "固定值"
                        IDCell.Offset(0, 4).Value = "加"
                    Case "万分比"
                        IDCell.Offset(0, 4).Value = "增加(万分比)"
                    Case Else
                        IDCell.Offset(0, 4).Value = "加"
                End Select
                IDCell.Offset(0, 5).Value = rowValue.Offset(0, 2)

            ElseIf rowValue = "PropResonance" Then
                IDCell.Offset(0, 2).Value = rowValue
                IDCell.Offset(0, 3).Value = rowValue.Offset(0, 1)
                IDCell.Offset(0, 4).Value = rowValue.Offset(0, 2)
                IDCell.Offset(0, 5).Value = rowValue.Offset(0, 3)
                IDCell.Offset(0, 6).Value = rowValue.Offset(0, 4)
                IDCell.Offset(0, 7).Value = rowValue.Offset(0, 5)

            ElseIf rowValue = "ModPropMaxORMin" Then
                IDCell.Offset(0, 2).Value = rowValue
                IDCell.Offset(0, 3).Value = rowValue.Offset(0, 1)
                IDCell.Offset(0, 4).Value = rowValue.Offset(0, 2)
                Select Case rowValue.Offset(0, 3).Value
                    Case "固定值"
                        IDCell.Offset(0, 5).Value = "0"
                    Case "万分比"
                        IDCell.Offset(0, 5).Value = "1"
                    Case Else
                        IDCell.Offset(0, 5).Value = "0"
                End Select

            ElseIf rowValue = "BuffDoHurt" Then
                IDCell.Offset(0, 2).Value = rowValue
                IDCell.Offset(0, 3).Value = rowValue.Offset(0, 1)
                Select Case rowValue.Offset(0, 2).Value
                    Case "固定值", "0"
                        IDCell.Offset(0, 4).Value = "0"
                    Case "当前生命万分比", "1"
                        IDCell.Offset(0, 4).Value = "1"
                    Case "最大生命万分比", "2"
                        IDCell.Offset(0, 4).Value = "2"
                    Case Else
                        IDCell.Offset(0, 4).Value = "0"
                End Select
                IDCell.Offset(0, 5).Value = rowValue.Offset(0, 3)

            ElseIf rowValue = "ModAttackDataProps" Then
                IDCell.Offset(0, 2).Value = rowValue
                IDCell.Offset(0, 3).Value = rowValue.Offset(0, 2)
                IDCell.Offset(0, 4).Value = "加"
                IDCell.Offset(0, 5).Value = rowValue.Offset(0, 3)
                IDCell.Offset(0, 6).Value = rowValue.Offset(0, 1)

            ElseIf rowValue = "ModAttackDataPropsByTag" Then
                IDCell.Offset(0, 2).Value = rowValue
                IDCell.Offset(0, 3).Value = rowValue.Offset(0, 2)
                IDCell.Offset(0, 4).Value = "加"
                IDCell.Offset(0, 5).Value = rowValue.Offset(0, 3)
                IDCell.Offset(0, 6).Value = rowValue.Offset(0, 1)

            Else
                IDCell.Offset(0, 2).Value = rowValue
                IDCell.Offset(0, 3).Value = rowValue.Offset(0, 1)
                IDCell.Offset(0, 4).Value = rowValue.Offset(0, 2)
                IDCell.Offset(0, 5).Value = rowValue.Offset(0, 3)
            End If

            IDCell.Offset(0, 8).Value = rowValue.Offset(0, 6)

            ' 将EffectIndex添加到buff表
            If EffectStr = "" Then
                EffectStr = (OffsetNumber + EffectIndex)
            Else
                EffectStr = EffectStr & "," & (OffsetNumber + EffectIndex)
            End If
            adwBuffEffect.Offset(i - tempRow + 2).Value = EffectStr
            EffectIndex = EffectIndex + 1
        End If
    Next i
End Function