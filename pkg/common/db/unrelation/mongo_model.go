package unrelation


// deleteMsgByLogic
//func (d *db.DataBases) DelMsgBySeqList(userID string, seqList []uint32, operationID string) (totalUnexistSeqList []uint32, err error) {
//	log.Debug(operationID, utils.GetSelfFuncName(), "args ", userID, seqList)
//	sortkeys.Uint32s(seqList)
//	suffixUserID2SubSeqList := func(uid string, seqList []uint32) map[string][]uint32 {
//		t := make(map[string][]uint32)
//		for i := 0; i < len(seqList); i++ {
//			seqUid := getSeqUid(uid, seqList[i])
//			if value, ok := t[seqUid]; !ok {
//				var temp []uint32
//				t[seqUid] = append(temp, seqList[i])
//			} else {
//				t[seqUid] = append(value, seqList[i])
//			}
//		}
//		return t
//	}(userID, seqList)
//
//	lock := sync.Mutex{}
//	var wg sync.WaitGroup
//	wg.Add(len(suffixUserID2SubSeqList))
//	for k, v := range suffixUserID2SubSeqList {
//		go func(suffixUserID string, subSeqList []uint32, operationID string) {
//			defer wg.Done()
//			unexistSeqList, err := d.DelMsgBySeqListInOneDoc(suffixUserID, subSeqList, operationID)
//			if err != nil {
//				log.Error(operationID, "DelMsgBySeqListInOneDoc failed ", err.Error(), suffixUserID, subSeqList)
//				return
//			}
//			lock.Lock()
//			totalUnexistSeqList = append(totalUnexistSeqList, unexistSeqList...)
//			lock.Unlock()
//		}(k, v, operationID)
//	}
//	return totalUnexistSeqList, err
//}
//
//func (d *db.DataBases) DelMsgBySeqListInOneDoc(suffixUserID string, seqList []uint32, operationID string) ([]uint32, error) {
//	log.Debug(operationID, utils.GetSelfFuncName(), "args ", suffixUserID, seqList)
//	seqMsgList, indexList, unexistSeqList, err := d.GetMsgAndIndexBySeqListInOneMongo2(suffixUserID, seqList, operationID)
//	if err != nil {
//		return nil, utils.Wrap(err, "")
//	}
//	for i, v := range seqMsgList {
//		if err := d.ReplaceMsgByIndex(suffixUserID, v, operationID, indexList[i]); err != nil {
//			return nil, utils.Wrap(err, "")
//		}
//	}
//	return unexistSeqList, nil
//}

// deleteMsgByLogic
//func (d *db.DataBases) DelMsgLogic(uid string, seqList []uint32, operationID string) error {
//	sortkeys.Uint32s(seqList)
//	seqMsgs, err := d.GetMsgBySeqListMongo2(uid, seqList, operationID)
//	if err != nil {
//		return utils.Wrap(err, "")
//	}
//	for _, seqMsg := range seqMsgs {
//		log.NewDebug(operationID, utils.GetSelfFuncName(), *seqMsg)
//		seqMsg.Status = constant.MsgDeleted
//		if err = d.ReplaceMsgBySeq(uid, seqMsg, operationID); err != nil {
//			log.NewError(operationID, utils.GetSelfFuncName(), "ReplaceMsgListBySeq error", err.Error())
//		}
//	}
//	return nil
//}

//func (d *db.DataBases) ReplaceMsgByIndex(suffixUserID string, msg *open_im_sdk.MsgData, operationID string, seqIndex int) error {
//	log.NewInfo(operationID, utils.GetSelfFuncName(), suffixUserID, *msg)
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	s := fmt.Sprintf("msg.%d.msg", seqIndex)
//	log.NewDebug(operationID, utils.GetSelfFuncName(), seqIndex, s)
//	msg.Status = constant.MsgDeleted
//	bytes, err := proto.Marshal(msg)
//	if err != nil {
//		log.NewError(operationID, utils.GetSelfFuncName(), "proto marshal failed ", err.Error(), msg.String())
//		return utils.Wrap(err, "")
//	}
//	updateResult, err := c.UpdateOne(ctx, bson.M{"uid": suffixUserID}, bson.M{"$set": bson.M{s: bytes}})
//	log.NewInfo(operationID, utils.GetSelfFuncName(), updateResult)
//	if err != nil {
//		log.NewError(operationID, utils.GetSelfFuncName(), "UpdateOne", err.Error())
//		return utils.Wrap(err, "")
//	}
//	return nil
//}

//func (d *db.DataBases) ReplaceMsgBySeq(uid string, msg *open_im_sdk.MsgData, operationID string) error {
//	log.NewInfo(operationID, utils.GetSelfFuncName(), uid, *msg)
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	uid = getSeqUid(uid, msg.Seq)
//	seqIndex := getMsgIndex(msg.Seq)
//	s := fmt.Sprintf("msg.%d.msg", seqIndex)
//	log.NewDebug(operationID, utils.GetSelfFuncName(), seqIndex, s)
//	bytes, err := proto.Marshal(msg)
//	if err != nil {
//		log.NewError(operationID, utils.GetSelfFuncName(), "proto marshal", err.Error())
//		return utils.Wrap(err, "")
//	}
//
//	updateResult, err := c.UpdateOne(
//		ctx, bson.M{"uid": uid},
//		bson.M{"$set": bson.M{s: bytes}})
//	log.NewInfo(operationID, utils.GetSelfFuncName(), updateResult)
//	if err != nil {
//		log.NewError(operationID, utils.GetSelfFuncName(), "UpdateOne", err.Error())
//		return utils.Wrap(err, "")
//	}
//	return nil
//}
//
//func (d *db.DataBases) UpdateOneMsgList(msg *UserChat) error {
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	_, err := c.UpdateOne(ctx, bson.M{"uid": msg.UID}, bson.M{"$set": bson.M{"msg": msg.Msg}})
//	return err
//}
//
//func (d *db.DataBases) GetMsgBySeqList(uid string, seqList []uint32, operationID string) (seqMsg []*open_im_sdk.MsgData, err error) {
//	log.NewInfo(operationID, utils.GetSelfFuncName(), uid, seqList)
//	var hasSeqList []uint32
//	singleCount := 0
//	session := d.mgoSession.Clone()
//	if session == nil {
//		return nil, errors.New("session == nil")
//	}
//	defer session.Close()
//	c := session.DB(config.Config.Mongo.DBDatabase).C(cChat)
//	m := func(uid string, seqList []uint32) map[string][]uint32 {
//		t := make(map[string][]uint32)
//		for i := 0; i < len(seqList); i++ {
//			seqUid := getSeqUid(uid, seqList[i])
//			if value, ok := t[seqUid]; !ok {
//				var temp []uint32
//				t[seqUid] = append(temp, seqList[i])
//			} else {
//				t[seqUid] = append(value, seqList[i])
//			}
//		}
//		return t
//	}(uid, seqList)
//	sChat := UserChat{}
//	for seqUid, value := range m {
//		if err = c.Find(bson.M{"uid": seqUid}).One(&sChat); err != nil {
//			log.NewError(operationID, "not find seqUid", seqUid, value, uid, seqList, err.Error())
//			continue
//		}
//		singleCount = 0
//		for i := 0; i < len(sChat.Msg); i++ {
//			msg := new(open_im_sdk.MsgData)
//			if err = proto.Unmarshal(sChat.Msg[i].Msg, msg); err != nil {
//				log.NewError(operationID, "Unmarshal err", seqUid, value, uid, seqList, err.Error())
//				return nil, err
//			}
//			if isContainInt32(msg.Seq, value) {
//				seqMsg = append(seqMsg, msg)
//				hasSeqList = append(hasSeqList, msg.Seq)
//				singleCount++
//				if singleCount == len(value) {
//					break
//				}
//			}
//		}
//	}
//	if len(hasSeqList) != len(seqList) {
//		var diff []uint32
//		diff = utils.Difference(hasSeqList, seqList)
//		exceptionMSg := genExceptionMessageBySeqList(diff)
//		seqMsg = append(seqMsg, exceptionMSg...)
//
//	}
//	return seqMsg, nil
//}
//
//func (d *db.DataBases) GetUserMsgListByIndex(ID string, index int64) (*UserChat, error) {
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	regex := fmt.Sprintf("^%s", ID)
//	findOpts := options.Find().SetLimit(1).SetSkip(index).SetSort(bson.M{"uid": 1})
//	var msgs []UserChat
//	//primitive.Regex{Pattern: regex}
//	cursor, err := c.Find(ctx, bson.M{"uid": primitive.Regex{Pattern: regex}}, findOpts)
//	if err != nil {
//		return nil, utils.Wrap(err, "")
//	}
//	err = cursor.All(context.Background(), &msgs)
//	if err != nil {
//		return nil, utils.Wrap(err, fmt.Sprintf("cursor is %s", cursor.Current.String()))
//	}
//	if len(msgs) > 0 {
//		return &msgs[0], nil
//	} else {
//		return nil, ErrMsgListNotExist
//	}
//}
//
//func (d *db.DataBases) DelMongoMsgs(IDList []string) error {
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	_, err := c.DeleteMany(ctx, bson.M{"uid": bson.M{"$in": IDList}})
//	return err
//}
//
//func (d *db.DataBases) ReplaceMsgToBlankByIndex(suffixID string, index int) (replaceMaxSeq uint32, err error) {
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	userChat := &UserChat{}
//	err = c.FindOne(ctx, bson.M{"uid": suffixID}).Decode(&userChat)
//	if err != nil {
//		return 0, err
//	}
//	for i, msg := range userChat.Msg {
//		if i <= index {
//			msgPb := &open_im_sdk.MsgData{}
//			if err = proto.Unmarshal(msg.Msg, msgPb); err != nil {
//				continue
//			}
//			newMsgPb := &open_im_sdk.MsgData{Seq: msgPb.Seq}
//			bytes, err := proto.Marshal(newMsgPb)
//			if err != nil {
//				continue
//			}
//			msg.Msg = bytes
//			msg.SendTime = 0
//			replaceMaxSeq = msgPb.Seq
//		}
//	}
//	_, err = c.UpdateOne(ctx, bson.M{"uid": suffixID}, bson.M{"$set": bson.M{"msg": userChat.Msg}})
//	return replaceMaxSeq, err
//}
//
//func (d *db.DataBases) GetNewestMsg(ID string) (msg *open_im_sdk.MsgData, err error) {
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	regex := fmt.Sprintf("^%s", ID)
//	findOpts := options.Find().SetLimit(1).SetSort(bson.M{"uid": -1})
//	var userChats []UserChat
//	cursor, err := c.Find(ctx, bson.M{"uid": bson.M{"$regex": regex}}, findOpts)
//	if err != nil {
//		return nil, err
//	}
//	err = cursor.All(ctx, &userChats)
//	if err != nil {
//		return nil, utils.Wrap(err, "")
//	}
//	if len(userChats) > 0 {
//		if len(userChats[0].Msg) > 0 {
//			msgPb := &open_im_sdk.MsgData{}
//			err = proto.Unmarshal(userChats[0].Msg[len(userChats[0].Msg)-1].Msg, msgPb)
//			if err != nil {
//				return nil, utils.Wrap(err, "")
//			}
//			return msgPb, nil
//		}
//		return nil, errors.New("len(userChats[0].Msg) < 0")
//	}
//	return nil, nil
//}
//
//func (d *db.DataBases) GetOldestMsg(ID string) (msg *open_im_sdk.MsgData, err error) {
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	regex := fmt.Sprintf("^%s", ID)
//	findOpts := options.Find().SetLimit(1).SetSort(bson.M{"uid": 1})
//	var userChats []UserChat
//	cursor, err := c.Find(ctx, bson.M{"uid": bson.M{"$regex": regex}}, findOpts)
//	if err != nil {
//		return nil, err
//	}
//	err = cursor.All(ctx, &userChats)
//	if err != nil {
//		return nil, utils.Wrap(err, "")
//	}
//	var oldestMsg []byte
//	if len(userChats) > 0 {
//		for _, v := range userChats[0].Msg {
//			if v.SendTime != 0 {
//				oldestMsg = v.Msg
//				break
//			}
//		}
//		if len(oldestMsg) == 0 {
//			oldestMsg = userChats[0].Msg[len(userChats[0].Msg)-1].Msg
//		}
//		msgPb := &open_im_sdk.MsgData{}
//		err = proto.Unmarshal(oldestMsg, msgPb)
//		if err != nil {
//			return nil, utils.Wrap(err, "")
//		}
//		return msgPb, nil
//	}
//	return nil, nil
//}
//
//func (d *db.DataBases) GetMsgBySeqListMongo2(uid string, seqList []uint32, operationID string) (seqMsg []*open_im_sdk.MsgData, err error) {
//	var hasSeqList []uint32
//	singleCount := 0
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//
//	m := func(uid string, seqList []uint32) map[string][]uint32 {
//		t := make(map[string][]uint32)
//		for i := 0; i < len(seqList); i++ {
//			seqUid := getSeqUid(uid, seqList[i])
//			if value, ok := t[seqUid]; !ok {
//				var temp []uint32
//				t[seqUid] = append(temp, seqList[i])
//			} else {
//				t[seqUid] = append(value, seqList[i])
//			}
//		}
//		return t
//	}(uid, seqList)
//	sChat := UserChat{}
//	for seqUid, value := range m {
//		if err = c.FindOne(ctx, bson.M{"uid": seqUid}).Decode(&sChat); err != nil {
//			log.NewError(operationID, "not find seqUid", seqUid, value, uid, seqList, err.Error())
//			continue
//		}
//		singleCount = 0
//		for i := 0; i < len(sChat.Msg); i++ {
//			msg := new(open_im_sdk.MsgData)
//			if err = proto.Unmarshal(sChat.Msg[i].Msg, msg); err != nil {
//				log.NewError(operationID, "Unmarshal err", seqUid, value, uid, seqList, err.Error())
//				return nil, err
//			}
//			if isContainInt32(msg.Seq, value) {
//				seqMsg = append(seqMsg, msg)
//				hasSeqList = append(hasSeqList, msg.Seq)
//				singleCount++
//				if singleCount == len(value) {
//					break
//				}
//			}
//		}
//	}
//	if len(hasSeqList) != len(seqList) {
//		var diff []uint32
//		diff = utils.Difference(hasSeqList, seqList)
//		exceptionMSg := genExceptionMessageBySeqList(diff)
//		seqMsg = append(seqMsg, exceptionMSg...)
//
//	}
//	return seqMsg, nil
//}
//func (d *db.DataBases) GetSuperGroupMsgBySeqListMongo(groupID string, seqList []uint32, operationID string) (seqMsg []*open_im_sdk.MsgData, err error) {
//	var hasSeqList []uint32
//	singleCount := 0
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//
//	m := func(uid string, seqList []uint32) map[string][]uint32 {
//		t := make(map[string][]uint32)
//		for i := 0; i < len(seqList); i++ {
//			seqUid := getSeqUid(uid, seqList[i])
//			if value, ok := t[seqUid]; !ok {
//				var temp []uint32
//				t[seqUid] = append(temp, seqList[i])
//			} else {
//				t[seqUid] = append(value, seqList[i])
//			}
//		}
//		return t
//	}(groupID, seqList)
//	sChat := UserChat{}
//	for seqUid, value := range m {
//		if err = c.FindOne(ctx, bson.M{"uid": seqUid}).Decode(&sChat); err != nil {
//			log.NewError(operationID, "not find seqGroupID", seqUid, value, groupID, seqList, err.Error())
//			continue
//		}
//		singleCount = 0
//		for i := 0; i < len(sChat.Msg); i++ {
//			msg := new(open_im_sdk.MsgData)
//			if err = proto.Unmarshal(sChat.Msg[i].Msg, msg); err != nil {
//				log.NewError(operationID, "Unmarshal err", seqUid, value, groupID, seqList, err.Error())
//				return nil, err
//			}
//			if isContainInt32(msg.Seq, value) {
//				seqMsg = append(seqMsg, msg)
//				hasSeqList = append(hasSeqList, msg.Seq)
//				singleCount++
//				if singleCount == len(value) {
//					break
//				}
//			}
//		}
//	}
//	if len(hasSeqList) != len(seqList) {
//		var diff []uint32
//		diff = utils.Difference(hasSeqList, seqList)
//		exceptionMSg := genExceptionSuperGroupMessageBySeqList(diff, groupID)
//		seqMsg = append(seqMsg, exceptionMSg...)
//
//	}
//	return seqMsg, nil
//}
//
//func (d *db.DataBases) GetMsgAndIndexBySeqListInOneMongo2(suffixUserID string, seqList []uint32, operationID string) (seqMsg []*open_im_sdk.MsgData, indexList []int, unexistSeqList []uint32, err error) {
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	sChat := UserChat{}
//	if err = c.FindOne(ctx, bson.M{"uid": suffixUserID}).Decode(&sChat); err != nil {
//		log.NewError(operationID, "not find seqUid", suffixUserID, err.Error())
//		return nil, nil, nil, utils.Wrap(err, "")
//	}
//	singleCount := 0
//	var hasSeqList []uint32
//	for i := 0; i < len(sChat.Msg); i++ {
//		msg := new(open_im_sdk.MsgData)
//		if err = proto.Unmarshal(sChat.Msg[i].Msg, msg); err != nil {
//			log.NewError(operationID, "Unmarshal err", msg.String(), err.Error())
//			return nil, nil, nil, err
//		}
//		if isContainInt32(msg.Seq, seqList) {
//			indexList = append(indexList, i)
//			seqMsg = append(seqMsg, msg)
//			hasSeqList = append(hasSeqList, msg.Seq)
//			singleCount++
//			if singleCount == len(seqList) {
//				break
//			}
//		}
//	}
//	for _, i := range seqList {
//		if isContainInt32(i, hasSeqList) {
//			continue
//		}
//		unexistSeqList = append(unexistSeqList, i)
//	}
//	return seqMsg, indexList, unexistSeqList, nil
//}
//
//func genExceptionMessageBySeqList(seqList []uint32) (exceptionMsg []*open_im_sdk.MsgData) {
//	for _, v := range seqList {
//		msg := new(open_im_sdk.MsgData)
//		msg.Seq = v
//		exceptionMsg = append(exceptionMsg, msg)
//	}
//	return exceptionMsg
//}
//
//func genExceptionSuperGroupMessageBySeqList(seqList []uint32, groupID string) (exceptionMsg []*open_im_sdk.MsgData) {
//	for _, v := range seqList {
//		msg := new(open_im_sdk.MsgData)
//		msg.Seq = v
//		msg.GroupID = groupID
//		msg.SessionType = constant.SuperGroupChatType
//		exceptionMsg = append(exceptionMsg, msg)
//	}
//	return exceptionMsg
//}
//
//func (d *db.DataBases) SaveUserChatMongo2(uid string, sendTime int64, m *pbMsg.MsgDataToDB) error {
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	newTime := getCurrentTimestampByMill()
//	operationID := ""
//	seqUid := getSeqUid(uid, m.MsgData.Seq)
//	filter := bson.M{"uid": seqUid}
//	var err error
//	sMsg := MsgInfo{}
//	sMsg.SendTime = sendTime
//	if sMsg.Msg, err = proto.Marshal(m.MsgData); err != nil {
//		return utils.Wrap(err, "")
//	}
//	err = c.FindOneAndUpdate(ctx, filter, bson.M{"$push": bson.M{"msg": sMsg}}).Err()
//	log.NewWarn(operationID, "get mgoSession cost time", getCurrentTimestampByMill()-newTime)
//	if err != nil {
//		sChat := UserChat{}
//		sChat.UID = seqUid
//		sChat.Msg = append(sChat.Msg, sMsg)
//		if _, err = c.InsertOne(ctx, &sChat); err != nil {
//			log.NewDebug(operationID, "InsertOne failed", filter)
//			return utils.Wrap(err, "")
//		}
//	} else {
//		log.NewDebug(operationID, "FindOneAndUpdate ok", filter)
//	}
//
//	log.NewDebug(operationID, "find mgo uid cost time", getCurrentTimestampByMill()-newTime)
//	return nil
//}

//
//func (d *DataBases) SaveUserChatListMongo2(uid string, sendTime int64, msgList []*pbMsg.MsgDataToDB) error {
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	newTime := getCurrentTimestampByMill()
//	operationID := ""
//	seqUid := ""
//	msgListToMongo := make([]MsgInfo, 0)
//
//	for _, m := range msgList {
//		seqUid = getSeqUid(uid, m.MsgData.Seq)
//		var err error
//		sMsg := MsgInfo{}
//		sMsg.SendTime = sendTime
//		if sMsg.Msg, err = proto.Marshal(m.MsgData); err != nil {
//			return utils.Wrap(err, "")
//		}
//		msgListToMongo = append(msgListToMongo, sMsg)
//	}
//
//	filter := bson.M{"uid": seqUid}
//	log.NewDebug(operationID, "filter ", seqUid)
//	err := c.FindOneAndUpdate(ctx, filter, bson.M{"$push": bson.M{"msg": bson.M{"$each": msgListToMongo}}}).Err()
//	log.NewWarn(operationID, "get mgoSession cost time", getCurrentTimestampByMill()-newTime)
//	if err != nil {
//		sChat := UserChat{}
//		sChat.UID = seqUid
//		sChat.Msg = msgListToMongo
//
//		if _, err = c.InsertOne(ctx, &sChat); err != nil {
//			log.NewError(operationID, "InsertOne failed", filter, err.Error(), sChat)
//			return utils.Wrap(err, "")
//		}
//	} else {
//		log.NewDebug(operationID, "FindOneAndUpdate ok", filter)
//	}
//
//	log.NewDebug(operationID, "find mgo uid cost time", getCurrentTimestampByMill()-newTime)
//	return nil
//}

//func (d *db.DataBases) SaveUserChat(uid string, sendTime int64, m *pbMsg.MsgDataToDB) error {
//	var seqUid string
//	newTime := getCurrentTimestampByMill()
//	session := d.mgoSession.Clone()
//	if session == nil {
//		return errors.New("session == nil")
//	}
//	defer session.Close()
//	log.NewDebug("", "get mgoSession cost time", getCurrentTimestampByMill()-newTime)
//	c := session.DB(config.Config.Mongo.DBDatabase).C(cChat)
//	seqUid = getSeqUid(uid, m.MsgData.Seq)
//	n, err := c.Find(bson.M{"uid": seqUid}).Count()
//	if err != nil {
//		return err
//	}
//	log.NewDebug("", "find mgo uid cost time", getCurrentTimestampByMill()-newTime)
//	sMsg := MsgInfo{}
//	sMsg.SendTime = sendTime
//	if sMsg.Msg, err = proto.Marshal(m.MsgData); err != nil {
//		return err
//	}
//	if n == 0 {
//		sChat := UserChat{}
//		sChat.UID = seqUid
//		sChat.Msg = append(sChat.Msg, sMsg)
//		err = c.Insert(&sChat)
//		if err != nil {
//			return err
//		}
//	} else {
//		err = c.Update(bson.M{"uid": seqUid}, bson.M{"$push": bson.M{"msg": sMsg}})
//		if err != nil {
//			return err
//		}
//	}
//	log.NewDebug("", "insert mgo data cost time", getCurrentTimestampByMill()-newTime)
//	return nil
//}
//
//func (d *db.DataBases) DelUserChatMongo2(uid string) error {
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
//	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(cChat)
//	filter := bson.M{"uid": uid}
//
//	delTime := time.Now().Unix() - int64(config.Config.Mongo.DBRetainChatRecords)*24*3600
//	if _, err := c.UpdateOne(ctx, filter, bson.M{"$pull": bson.M{"msg": bson.M{"sendtime": bson.M{"$lte": delTime}}}}); err != nil {
//		return utils.Wrap(err, "")
//	}
//	return nil
//}
//
//func (d *db.DataBases) MgoSkipUID(count int) (string, error) {
//	return "", nil
	//session := d.mgoSession.Clone()
	//if session == nil {
	//	return "", errors.New("session == nil")
	//}
	//defer session.Close()
	//
	//c := session.DB(config.Config.Mongo.DBDatabase).C(cChat)
	//
	//sChat := UserChat{}
	//c.Find(nil).Skip(count).Limit(1).One(&sChat)
	//return sChat.UID, nil
}

//func generateTagID(tagName, userID string) string {
//	return utils.Md5(tagName + userID + strconv.Itoa(rand.Int()) + time.Now().String())
//}

//func generateWorkMomentID(userID string) string {
//	return utils.Md5(userID + strconv.Itoa(rand.Int()) + time.Now().String())
//}

//func generateWorkMomentCommentID(workMomentID string) string {
//	return utils.Md5(workMomentID + strconv.Itoa(rand.Int()) + time.Now().String())
//}

//func getCurrentTimestampByMill() int64 {
//	return time.Now().UnixNano() / 1e6
//}
//func GetCurrentTimestampByMill() int64 {
//	return time.Now().UnixNano() / 1e6
//}

//func getSeqUid(uid string, seq uint32) string {
//	seqSuffix := seq / singleGocMsgNum
//	return indexGen(uid, seqSuffix)
//}
//
//func getSeqUserIDList(userID string, maxSeq uint32) []string {
//	seqMaxSuffix := maxSeq / singleGocMsgNum
//	var seqUserIDList []string
//	for i := 0; i <= int(seqMaxSuffix); i++ {
//		seqUserID := indexGen(userID, uint32(i))
//		seqUserIDList = append(seqUserIDList, seqUserID)
//	}
//	return seqUserIDList
//}
//
//func getSeqSuperGroupID(groupID string, seq uint32) string {
//	seqSuffix := seq / singleGocMsgNum
//	return superGroupIndexGen(groupID, seqSuffix)
//}
//
//func GetSeqUid(uid string, seq uint32) string {
//	return getSeqUid(uid, seq)
//}
//
//func getMsgIndex(seq uint32) int {
//	seqSuffix := seq / singleGocMsgNum
//	var index uint32
//	if seqSuffix == 0 {
//		index = (seq - seqSuffix*singleGocMsgNum) - 1
//	} else {
//		index = seq - seqSuffix*singleGocMsgNum
//	}
//	return int(index)
//}

//func isContainInt32(target uint32, List []uint32) bool {
//	for _, element := range List {
//		if target == element {
//			return true
//		}
//	}
//	return false
//}
//
//func isNotContainInt32(target uint32, List []uint32) bool {
//	for _, i := range List {
//		if i == target {
//			return false
//		}
//	}
//	return true
//}
//
//func indexGen(uid string, seqSuffix uint32) string {
//	return uid + ":" + strconv.FormatInt(int64(seqSuffix), 10)
//}

//func superGroupIndexGen(groupID string, seqSuffix uint32) string {
//	return "super_group_" + groupID + ":" + strconv.FormatInt(int64(seqSuffix), 10)
//}
