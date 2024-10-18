from epyxid import XID, xid_create

def new():
     _xid: XID = xid_create()
     return _xid.to_str()