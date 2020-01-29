package main

func downloadBaseInfos(r *BaseInfoResponse) {
	dlBaseInfo(r.MsgBody.BaseInfoItem.AreaDownloadURL)
	dlBaseInfo(r.MsgBody.BaseInfoItem.RouteDownloadURL)
	dlBaseInfo(r.MsgBody.BaseInfoItem.RouteLineDownloadURL)
	dlBaseInfo(r.MsgBody.BaseInfoItem.RouteStationDownloadURL)
	dlBaseInfo(r.MsgBody.BaseInfoItem.StationDownloadURL)
}
