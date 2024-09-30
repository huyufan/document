package geecache

import pb "exec/go/exec/cache/geecache/geecachepb"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	//Get(group string, key string) ([]byte, error)
	Get(in *pb.Request, out *pb.Response) error
}
