from typing import Union
import types
_mock_data={}
class Pv:
	def get(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pv.get", param0))
		if "get" not in _mock_data['pv']:
			return None
		node = _mock_data['pv']['get']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def getw(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pv.getw", param0))
		if "getw" not in _mock_data['pv']:
			return None
		node = _mock_data['pv']['getw']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def gete(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pv.gete", param0))
		if "gete" not in _mock_data['pv']:
			return None
		node = _mock_data['pv']['gete']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def getvn(self, param0: str, param1: int) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pv.getvn", param0, param1))
		if "getvn" not in _mock_data['pv']:
			return None
		node = _mock_data['pv']['getvn']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def getvs(self, param0: str, param1: str) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pv.getvs", param0, param1))
		if "getvs" not in _mock_data['pv']:
			return None
		node = _mock_data['pv']['getvs']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def seti(self, param0: str, param1: int) -> bool:
		print("Calling %s, %s, %s" % ("pv.seti", param0, param1))
		if "seti" not in _mock_data['pv']:
			return True
		node = _mock_data['pv']['seti']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return True
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return True
		return node

	def sets(self, param0: str, param1: str) -> bool:
		print("Calling %s, %s, %s" % ("pv.sets", param0, param1))
		if "sets" not in _mock_data['pv']:
			return True
		node = _mock_data['pv']['sets']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return True
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return True
		return node

	def unset(self, param0: str) -> bool:
		print("Calling %s, %s" % ("pv.unset", param0))
		if "unset" not in _mock_data['pv']:
			return True
		node = _mock_data['pv']['unset']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return True
		return node

	def is_null(self, param0: str) -> bool:
		print("Calling %s, %s" % ("pv.is_null", param0))
		if "is_null" not in _mock_data['pv']:
			return True
		node = _mock_data['pv']['is_null']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return True
		return node

class Hdr:
	def append(self, param0: str) -> int:
		print("Calling %s, %s" % ("hdr.append", param0))
		if "append" not in _mock_data['hdr']:
			return 1
		node = _mock_data['hdr']['append']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def append_after(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("hdr.append_after", param0, param1))
		if "append_after" not in _mock_data['hdr']:
			return 1
		node = _mock_data['hdr']['append_after']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def insert(self, param0: str) -> int:
		print("Calling %s, %s" % ("hdr.insert", param0))
		if "insert" not in _mock_data['hdr']:
			return 1
		node = _mock_data['hdr']['insert']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def insert_before(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("hdr.insert_before", param0, param1))
		if "insert_before" not in _mock_data['hdr']:
			return 1
		node = _mock_data['hdr']['insert_before']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def remove(self, param0: str) -> int:
		print("Calling %s, %s" % ("hdr.remove", param0))
		if "remove" not in _mock_data['hdr']:
			return 1
		node = _mock_data['hdr']['remove']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def rmappend(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("hdr.rmappend", param0, param1))
		if "rmappend" not in _mock_data['hdr']:
			return 1
		node = _mock_data['hdr']['rmappend']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def rminsert(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("hdr.rminsert", param0, param1))
		if "rminsert" not in _mock_data['hdr']:
			return 1
		node = _mock_data['hdr']['rminsert']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def is_present(self, param0: str) -> int:
		print("Calling %s, %s" % ("hdr.is_present", param0))
		if "is_present" not in _mock_data['hdr']:
			return 1
		node = _mock_data['hdr']['is_present']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def append_to_reply(self, param0: str) -> int:
		print("Calling %s, %s" % ("hdr.append_to_reply", param0))
		if "append_to_reply" not in _mock_data['hdr']:
			return 1
		node = _mock_data['hdr']['append_to_reply']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def get(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("hdr.get", param0))
		if "get" not in _mock_data['hdr']:
			return None
		node = _mock_data['hdr']['get']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def gete(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("hdr.gete", param0))
		if "gete" not in _mock_data['hdr']:
			return None
		node = _mock_data['hdr']['gete']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def getw(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("hdr.getw", param0))
		if "getw" not in _mock_data['hdr']:
			return None
		node = _mock_data['hdr']['getw']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def get_idx(self, param0: str, param1: int) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("hdr.get_idx", param0, param1))
		if "get_idx" not in _mock_data['hdr']:
			return None
		node = _mock_data['hdr']['get_idx']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def gete_idx(self, param0: str, param1: int) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("hdr.gete_idx", param0, param1))
		if "gete_idx" not in _mock_data['hdr']:
			return None
		node = _mock_data['hdr']['gete_idx']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def getw_idx(self, param0: str, param1: int) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("hdr.getw_idx", param0, param1))
		if "getw_idx" not in _mock_data['hdr']:
			return None
		node = _mock_data['hdr']['getw_idx']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def match_content(self, param0: str, param1: str, param2: str, param3: str) -> bool:
		print("Calling %s, %s, %s, %s, %s" % ("hdr.match_content", param0, param1, param2, param3))
		if "match_content" not in _mock_data['hdr']:
			return True
		node = _mock_data['hdr']['match_content']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2, param3)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return True
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return True
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return True
		if not isinstance(node, dict):
			return node
		if str(param3) in node:
			node = node[str(param3)]
		else:
			return True
		return node

class Tm:
	def t_relay(self) -> int:
		print("Calling %s" % ("tm.t_relay"))
		if "t_relay" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_relay']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_on_branch(self, param0: str) -> int:
		print("Calling %s, %s" % ("tm.t_on_branch", param0))
		if "t_on_branch" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_on_branch']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_on_failure(self, param0: str) -> int:
		print("Calling %s, %s" % ("tm.t_on_failure", param0))
		if "t_on_failure" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_on_failure']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_on_branch_failure(self, param0: str) -> int:
		print("Calling %s, %s" % ("tm.t_on_branch_failure", param0))
		if "t_on_branch_failure" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_on_branch_failure']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_on_reply(self, param0: str) -> int:
		print("Calling %s, %s" % ("tm.t_on_reply", param0))
		if "t_on_reply" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_on_reply']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_reply(self, param0: int, param1: str) -> int:
		print("Calling %s, %s, %s" % ("tm.t_reply", param0, param1))
		if "t_reply" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_reply']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def t_send_reply(self, param0: int, param1: str) -> int:
		print("Calling %s, %s, %s" % ("tm.t_send_reply", param0, param1))
		if "t_send_reply" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_send_reply']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def t_check_trans(self) -> int:
		print("Calling %s" % ("tm.t_check_trans"))
		if "t_check_trans" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_check_trans']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_is_canceled(self) -> int:
		print("Calling %s" % ("tm.t_is_canceled"))
		if "t_is_canceled" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_is_canceled']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_newtran(self) -> int:
		print("Calling %s" % ("tm.t_newtran"))
		if "t_newtran" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_newtran']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_release(self) -> int:
		print("Calling %s" % ("tm.t_release"))
		if "t_release" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_release']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_replicate(self, param0: str) -> int:
		print("Calling %s, %s" % ("tm.t_replicate", param0))
		if "t_replicate" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_replicate']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_is_set(self, param0: str) -> int:
		print("Calling %s, %s" % ("tm.t_is_set", param0))
		if "t_is_set" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_is_set']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_lookup_request(self) -> int:
		print("Calling %s" % ("tm.t_lookup_request"))
		if "t_lookup_request" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_lookup_request']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_lookup_cancel(self) -> int:
		print("Calling %s" % ("tm.t_lookup_cancel"))
		if "t_lookup_cancel" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_lookup_cancel']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_lookup_cancel_flags(self, param0: int) -> int:
		print("Calling %s, %s" % ("tm.t_lookup_cancel_flags", param0))
		if "t_lookup_cancel_flags" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_lookup_cancel_flags']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_retransmit_reply(self) -> int:
		print("Calling %s" % ("tm.t_retransmit_reply"))
		if "t_retransmit_reply" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_retransmit_reply']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_set_fr_inv(self, param0: int) -> int:
		print("Calling %s, %s" % ("tm.t_set_fr_inv", param0))
		if "t_set_fr_inv" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_set_fr_inv']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_set_fr(self, param0: int, param1: int) -> int:
		print("Calling %s, %s, %s" % ("tm.t_set_fr", param0, param1))
		if "t_set_fr" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_set_fr']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def t_reset_fr(self) -> int:
		print("Calling %s" % ("tm.t_reset_fr"))
		if "t_reset_fr" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_reset_fr']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_set_max_lifetime(self, param0: int, param1: int) -> int:
		print("Calling %s, %s, %s" % ("tm.t_set_max_lifetime", param0, param1))
		if "t_set_max_lifetime" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_set_max_lifetime']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def t_reset_max_lifetime(self) -> int:
		print("Calling %s" % ("tm.t_reset_max_lifetime"))
		if "t_reset_max_lifetime" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_reset_max_lifetime']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_set_retr(self, param0: int, param1: int) -> int:
		print("Calling %s, %s, %s" % ("tm.t_set_retr", param0, param1))
		if "t_set_retr" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_set_retr']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def t_reset_retr(self) -> int:
		print("Calling %s" % ("tm.t_reset_retr"))
		if "t_reset_retr" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_reset_retr']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_uac_send(self, param0: str, param1: str, param2: str, param3: str, param4: str, param5: str) -> int:
		print("Calling %s, %s, %s, %s, %s, %s, %s" % ("tm.t_uac_send", param0, param1, param2, param3, param4, param5))
		if "t_uac_send" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_uac_send']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2, param3, param4, param5)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param3) in node:
			node = node[str(param3)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param4) in node:
			node = node[str(param4)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param5) in node:
			node = node[str(param5)]
		else:
			return 1
		return node

	def t_load_contacts(self) -> int:
		print("Calling %s" % ("tm.t_load_contacts"))
		if "t_load_contacts" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_load_contacts']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def ki_t_load_contacts_mode(self, param0: int) -> int:
		print("Calling %s, %s" % ("tm.ki_t_load_contacts_mode", param0))
		if "ki_t_load_contacts_mode" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['ki_t_load_contacts_mode']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_next_contacts(self) -> int:
		print("Calling %s" % ("tm.t_next_contacts"))
		if "t_next_contacts" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_next_contacts']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_next_contact_flow(self) -> int:
		print("Calling %s" % ("tm.t_next_contact_flow"))
		if "t_next_contact_flow" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_next_contact_flow']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_drop_replies_all(self) -> int:
		print("Calling %s" % ("tm.t_drop_replies_all"))
		if "t_drop_replies_all" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_drop_replies_all']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_drop_replies(self, param0: str) -> int:
		print("Calling %s, %s" % ("tm.t_drop_replies", param0))
		if "t_drop_replies" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_drop_replies']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_use_uac_headers(self) -> int:
		print("Calling %s" % ("tm.t_use_uac_headers"))
		if "t_use_uac_headers" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_use_uac_headers']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_save_lumps(self) -> int:
		print("Calling %s" % ("tm.t_save_lumps"))
		if "t_save_lumps" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_save_lumps']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_is_expired(self) -> int:
		print("Calling %s" % ("tm.t_is_expired"))
		if "t_is_expired" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_is_expired']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_check_status(self, param0: str) -> int:
		print("Calling %s, %s" % ("tm.t_check_status", param0))
		if "t_check_status" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_check_status']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_grep_status(self, param0: int) -> int:
		print("Calling %s, %s" % ("tm.t_grep_status", param0))
		if "t_grep_status" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_grep_status']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_is_retr_async_reply(self) -> int:
		print("Calling %s" % ("tm.t_is_retr_async_reply"))
		if "t_is_retr_async_reply" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_is_retr_async_reply']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_any_replied(self) -> int:
		print("Calling %s" % ("tm.t_any_replied"))
		if "t_any_replied" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_any_replied']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_any_timeout(self) -> int:
		print("Calling %s" % ("tm.t_any_timeout"))
		if "t_any_timeout" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_any_timeout']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_branch_replied(self) -> int:
		print("Calling %s" % ("tm.t_branch_replied"))
		if "t_branch_replied" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_branch_replied']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_branch_timeout(self) -> int:
		print("Calling %s" % ("tm.t_branch_timeout"))
		if "t_branch_timeout" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_branch_timeout']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_set_auto_inv_100(self, param0: int) -> int:
		print("Calling %s, %s" % ("tm.t_set_auto_inv_100", param0))
		if "t_set_auto_inv_100" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_set_auto_inv_100']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_set_disable_6xx(self, param0: int) -> int:
		print("Calling %s, %s" % ("tm.t_set_disable_6xx", param0))
		if "t_set_disable_6xx" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_set_disable_6xx']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_set_disable_failover(self, param0: int) -> int:
		print("Calling %s, %s" % ("tm.t_set_disable_failover", param0))
		if "t_set_disable_failover" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_set_disable_failover']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_set_no_e2e_cancel_reason(self, param0: int) -> int:
		print("Calling %s, %s" % ("tm.t_set_no_e2e_cancel_reason", param0))
		if "t_set_no_e2e_cancel_reason" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_set_no_e2e_cancel_reason']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_set_disable_internal_reply(self, param0: int) -> int:
		print("Calling %s, %s" % ("tm.t_set_disable_internal_reply", param0))
		if "t_set_disable_internal_reply" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_set_disable_internal_reply']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_relay_to_proxy(self, param0: str) -> int:
		print("Calling %s, %s" % ("tm.t_relay_to_proxy", param0))
		if "t_relay_to_proxy" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_relay_to_proxy']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_relay_to_flags(self, param0: int) -> int:
		print("Calling %s, %s" % ("tm.t_relay_to_flags", param0))
		if "t_relay_to_flags" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_relay_to_flags']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_relay_to_proxy_flags(self, param0: str, param1: int) -> int:
		print("Calling %s, %s, %s" % ("tm.t_relay_to_proxy_flags", param0, param1))
		if "t_relay_to_proxy_flags" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_relay_to_proxy_flags']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def t_relay_to_proto(self, param0: str) -> int:
		print("Calling %s, %s" % ("tm.t_relay_to_proto", param0))
		if "t_relay_to_proto" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_relay_to_proto']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_relay_to_proto_addr(self, param0: str, param1: str, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("tm.t_relay_to_proto_addr", param0, param1, param2))
		if "t_relay_to_proto_addr" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_relay_to_proto_addr']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def t_get_status_code(self) -> int:
		print("Calling %s" % ("tm.t_get_status_code"))
		if "t_get_status_code" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_get_status_code']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_get_branch_index(self) -> int:
		print("Calling %s" % ("tm.t_get_branch_index"))
		if "t_get_branch_index" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_get_branch_index']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_clean(self) -> int:
		print("Calling %s" % ("tm.t_clean"))
		if "t_clean" not in _mock_data['tm']:
			return 1
		node = _mock_data['tm']['t_clean']
		if isinstance(node, types.FunctionType):
			return node()
		return node

class Tmx:
	def t_precheck_trans(self) -> int:
		print("Calling %s" % ("tmx.t_precheck_trans"))
		if "t_precheck_trans" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_precheck_trans']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_is_request_route(self) -> int:
		print("Calling %s" % ("tmx.t_is_request_route"))
		if "t_is_request_route" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_is_request_route']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_is_reply_route(self) -> int:
		print("Calling %s" % ("tmx.t_is_reply_route"))
		if "t_is_reply_route" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_is_reply_route']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_is_failure_route(self) -> int:
		print("Calling %s" % ("tmx.t_is_failure_route"))
		if "t_is_failure_route" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_is_failure_route']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_is_branch_route(self) -> int:
		print("Calling %s" % ("tmx.t_is_branch_route"))
		if "t_is_branch_route" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_is_branch_route']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_suspend(self) -> int:
		print("Calling %s" % ("tmx.t_suspend"))
		if "t_suspend" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_suspend']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_continue(self, param0: int, param1: int, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("tmx.t_continue", param0, param1, param2))
		if "t_continue" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_continue']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def t_flush_flags(self) -> int:
		print("Calling %s" % ("tmx.t_flush_flags"))
		if "t_flush_flags" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_flush_flags']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_flush_xflags(self) -> int:
		print("Calling %s" % ("tmx.t_flush_xflags"))
		if "t_flush_xflags" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_flush_xflags']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_cancel_branches(self, param0: str) -> int:
		print("Calling %s, %s" % ("tmx.t_cancel_branches", param0))
		if "t_cancel_branches" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_cancel_branches']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def t_reuse_branch(self) -> int:
		print("Calling %s" % ("tmx.t_reuse_branch"))
		if "t_reuse_branch" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_reuse_branch']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_cancel_callid(self, param0: str, param1: str, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("tmx.t_cancel_callid", param0, param1, param2))
		if "t_cancel_callid" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_cancel_callid']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def t_cancel_callid_reason(self, param0: str, param1: str, param2: int, param3: int) -> int:
		print("Calling %s, %s, %s, %s, %s" % ("tmx.t_cancel_callid_reason", param0, param1, param2, param3))
		if "t_cancel_callid_reason" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_cancel_callid_reason']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2, param3)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param3) in node:
			node = node[str(param3)]
		else:
			return 1
		return node

	def t_reply_callid(self, param0: str, param1: str, param2: int, param3: str) -> int:
		print("Calling %s, %s, %s, %s, %s" % ("tmx.t_reply_callid", param0, param1, param2, param3))
		if "t_reply_callid" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_reply_callid']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2, param3)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param3) in node:
			node = node[str(param3)]
		else:
			return 1
		return node

	def t_drop(self) -> int:
		print("Calling %s" % ("tmx.t_drop"))
		if "t_drop" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_drop']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def t_drop_rcode(self, param0: int) -> int:
		print("Calling %s, %s" % ("tmx.t_drop_rcode", param0))
		if "t_drop_rcode" not in _mock_data['tmx']:
			return 1
		node = _mock_data['tmx']['t_drop_rcode']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

class Sl:
	def sl_send_reply(self, param0: int, param1: str) -> int:
		print("Calling %s, %s, %s" % ("sl.sl_send_reply", param0, param1))
		if "sl_send_reply" not in _mock_data['sl']:
			return 1
		node = _mock_data['sl']['sl_send_reply']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def send_reply(self, param0: int, param1: str) -> int:
		print("Calling %s, %s, %s" % ("sl.send_reply", param0, param1))
		if "send_reply" not in _mock_data['sl']:
			return 1
		node = _mock_data['sl']['send_reply']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def send_reply_mode(self, param0: int, param1: str, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("sl.send_reply_mode", param0, param1, param2))
		if "send_reply_mode" not in _mock_data['sl']:
			return 1
		node = _mock_data['sl']['send_reply_mode']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def sl_reply_error(self) -> int:
		print("Calling %s" % ("sl.sl_reply_error"))
		if "sl_reply_error" not in _mock_data['sl']:
			return 1
		node = _mock_data['sl']['sl_reply_error']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def sl_forward_reply(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("sl.sl_forward_reply", param0, param1))
		if "sl_forward_reply" not in _mock_data['sl']:
			return 1
		node = _mock_data['sl']['sl_forward_reply']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

class Rr:
	def record_route(self) -> int:
		print("Calling %s" % ("rr.record_route"))
		if "record_route" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['record_route']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def record_route_params(self, param0: str) -> int:
		print("Calling %s, %s" % ("rr.record_route_params", param0))
		if "record_route_params" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['record_route_params']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def loose_route(self) -> int:
		print("Calling %s" % ("rr.loose_route"))
		if "loose_route" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['loose_route']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def loose_route_preloaded(self) -> int:
		print("Calling %s" % ("rr.loose_route_preloaded"))
		if "loose_route_preloaded" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['loose_route_preloaded']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def loose_route_mode(self, param0: int) -> int:
		print("Calling %s, %s" % ("rr.loose_route_mode", param0))
		if "loose_route_mode" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['loose_route_mode']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def remove_record_route(self) -> int:
		print("Calling %s" % ("rr.remove_record_route"))
		if "remove_record_route" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['remove_record_route']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def add_rr_param(self, param0: str) -> int:
		print("Calling %s, %s" % ("rr.add_rr_param", param0))
		if "add_rr_param" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['add_rr_param']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def check_route_param(self, param0: str) -> int:
		print("Calling %s, %s" % ("rr.check_route_param", param0))
		if "check_route_param" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['check_route_param']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def is_direction(self, param0: str) -> int:
		print("Calling %s, %s" % ("rr.is_direction", param0))
		if "is_direction" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['is_direction']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def record_route_preset_one(self, param0: str) -> int:
		print("Calling %s, %s" % ("rr.record_route_preset_one", param0))
		if "record_route_preset_one" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['record_route_preset_one']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def record_route_preset(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("rr.record_route_preset", param0, param1))
		if "record_route_preset" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['record_route_preset']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def record_route_advertised_address(self, param0: str) -> int:
		print("Calling %s, %s" % ("rr.record_route_advertised_address", param0))
		if "record_route_advertised_address" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['record_route_advertised_address']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def next_hop_route(self) -> int:
		print("Calling %s" % ("rr.next_hop_route"))
		if "next_hop_route" not in _mock_data['rr']:
			return 1
		node = _mock_data['rr']['next_hop_route']
		if isinstance(node, types.FunctionType):
			return node()
		return node

class Pvx:
	def sbranch_set_ruri(self) -> int:
		print("Calling %s" % ("pvx.sbranch_set_ruri"))
		if "sbranch_set_ruri" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['sbranch_set_ruri']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def sbranch_append(self) -> int:
		print("Calling %s" % ("pvx.sbranch_append"))
		if "sbranch_append" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['sbranch_append']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def sbranch_reset(self) -> int:
		print("Calling %s" % ("pvx.sbranch_reset"))
		if "sbranch_reset" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['sbranch_reset']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def var_seti(self, param0: str, param1: int) -> int:
		print("Calling %s, %s, %s" % ("pvx.var_seti", param0, param1))
		if "var_seti" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['var_seti']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def var_sets(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.var_sets", param0, param1))
		if "var_sets" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['var_sets']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def var_get(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.var_get", param0))
		if "var_get" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['var_get']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def shv_seti(self, param0: str, param1: int) -> int:
		print("Calling %s, %s, %s" % ("pvx.shv_seti", param0, param1))
		if "shv_seti" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['shv_seti']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def shv_sets(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.shv_sets", param0, param1))
		if "shv_sets" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['shv_sets']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def shv_get(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.shv_get", param0))
		if "shv_get" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['shv_get']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def pv_var_to_xavp(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.pv_var_to_xavp", param0, param1))
		if "pv_var_to_xavp" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['pv_var_to_xavp']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def pv_xavp_to_var(self, param0: str) -> int:
		print("Calling %s, %s" % ("pvx.pv_xavp_to_var", param0))
		if "pv_xavp_to_var" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['pv_xavp_to_var']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def pv_xavp_print(self) -> int:
		print("Calling %s" % ("pvx.pv_xavp_print"))
		if "pv_xavp_print" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['pv_xavp_print']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def pv_xavu_print(self) -> int:
		print("Calling %s" % ("pvx.pv_xavu_print"))
		if "pv_xavu_print" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['pv_xavu_print']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def pv_xavi_print(self) -> int:
		print("Calling %s" % ("pvx.pv_xavi_print"))
		if "pv_xavi_print" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['pv_xavi_print']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def xavp_params_explode(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavp_params_explode", param0, param1))
		if "xavp_params_explode" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_params_explode']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavp_params_implode(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavp_params_implode", param0, param1))
		if "xavp_params_implode" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_params_implode']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavp_slist_explode(self, param0: str, param1: str, param2: str, param3: str) -> int:
		print("Calling %s, %s, %s, %s, %s" % ("pvx.xavp_slist_explode", param0, param1, param2, param3))
		if "xavp_slist_explode" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_slist_explode']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2, param3)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param3) in node:
			node = node[str(param3)]
		else:
			return 1
		return node

	def xavp_seti(self, param0: str, param1: int) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavp_seti", param0, param1))
		if "xavp_seti" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_seti']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavp_sets(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavp_sets", param0, param1))
		if "xavp_sets" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_sets']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavp_get(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.xavp_get", param0))
		if "xavp_get" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavp_get']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def xavp_gete(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.xavp_gete", param0))
		if "xavp_gete" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavp_gete']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def xavp_getw(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.xavp_getw", param0))
		if "xavp_getw" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavp_getw']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def xavp_getd(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.xavp_getd", param0))
		if "xavp_getd" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavp_getd']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def xavp_getd_p1(self, param0: str, param1: int) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavp_getd_p1", param0, param1))
		if "xavp_getd_p1" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavp_getd_p1']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavp_get_keys(self, param0: str, param1: int) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavp_get_keys", param0, param1))
		if "xavp_get_keys" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavp_get_keys']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavp_rm(self, param0: str) -> int:
		print("Calling %s, %s" % ("pvx.xavp_rm", param0))
		if "xavp_rm" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_rm']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xavp_is_null(self, param0: str) -> int:
		print("Calling %s, %s" % ("pvx.xavp_is_null", param0))
		if "xavp_is_null" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_is_null']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xavp_child_seti(self, param0: str, param1: str, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("pvx.xavp_child_seti", param0, param1, param2))
		if "xavp_child_seti" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_child_seti']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def xavp_child_sets(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("pvx.xavp_child_sets", param0, param1, param2))
		if "xavp_child_sets" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_child_sets']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def xavp_child_rm(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavp_child_rm", param0, param1))
		if "xavp_child_rm" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_child_rm']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavp_child_is_null(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavp_child_is_null", param0, param1))
		if "xavp_child_is_null" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_child_is_null']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavp_child_get(self, param0: str, param1: str) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavp_child_get", param0, param1))
		if "xavp_child_get" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavp_child_get']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavp_child_gete(self, param0: str, param1: str) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavp_child_gete", param0, param1))
		if "xavp_child_gete" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavp_child_gete']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavp_child_getw(self, param0: str, param1: str) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavp_child_getw", param0, param1))
		if "xavp_child_getw" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavp_child_getw']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavu_seti(self, param0: str, param1: int) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavu_seti", param0, param1))
		if "xavu_seti" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavu_seti']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavu_sets(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavu_sets", param0, param1))
		if "xavu_sets" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavu_sets']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavu_get(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.xavu_get", param0))
		if "xavu_get" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavu_get']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def xavu_gete(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.xavu_gete", param0))
		if "xavu_gete" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavu_gete']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def xavu_getw(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.xavu_getw", param0))
		if "xavu_getw" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavu_getw']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def xavu_rm(self, param0: str) -> int:
		print("Calling %s, %s" % ("pvx.xavu_rm", param0))
		if "xavu_rm" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavu_rm']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xavu_is_null(self, param0: str) -> int:
		print("Calling %s, %s" % ("pvx.xavu_is_null", param0))
		if "xavu_is_null" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavu_is_null']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xavu_child_seti(self, param0: str, param1: str, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("pvx.xavu_child_seti", param0, param1, param2))
		if "xavu_child_seti" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavu_child_seti']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def xavu_child_sets(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("pvx.xavu_child_sets", param0, param1, param2))
		if "xavu_child_sets" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavu_child_sets']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def xavu_child_rm(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavu_child_rm", param0, param1))
		if "xavu_child_rm" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavu_child_rm']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavu_child_is_null(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavu_child_is_null", param0, param1))
		if "xavu_child_is_null" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavu_child_is_null']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavu_child_get(self, param0: str, param1: str) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavu_child_get", param0, param1))
		if "xavu_child_get" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavu_child_get']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavu_child_gete(self, param0: str, param1: str) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavu_child_gete", param0, param1))
		if "xavu_child_gete" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavu_child_gete']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavu_child_getw(self, param0: str, param1: str) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavu_child_getw", param0, param1))
		if "xavu_child_getw" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavu_child_getw']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavi_seti(self, param0: str, param1: int) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavi_seti", param0, param1))
		if "xavi_seti" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavi_seti']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavi_sets(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavi_sets", param0, param1))
		if "xavi_sets" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavi_sets']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavi_get(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.xavi_get", param0))
		if "xavi_get" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavi_get']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def xavi_gete(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.xavi_gete", param0))
		if "xavi_gete" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavi_gete']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def xavi_getw(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.xavi_getw", param0))
		if "xavi_getw" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavi_getw']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def xavi_getd(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.xavi_getd", param0))
		if "xavi_getd" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavi_getd']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def xavi_getd_p1(self, param0: str, param1: int) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavi_getd_p1", param0, param1))
		if "xavi_getd_p1" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavi_getd_p1']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavi_get_keys(self, param0: str, param1: int) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavi_get_keys", param0, param1))
		if "xavi_get_keys" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavi_get_keys']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavi_rm(self, param0: str) -> int:
		print("Calling %s, %s" % ("pvx.xavi_rm", param0))
		if "xavi_rm" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavi_rm']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xavi_is_null(self, param0: str) -> int:
		print("Calling %s, %s" % ("pvx.xavi_is_null", param0))
		if "xavi_is_null" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavi_is_null']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xavi_child_seti(self, param0: str, param1: str, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("pvx.xavi_child_seti", param0, param1, param2))
		if "xavi_child_seti" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavi_child_seti']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def xavi_child_sets(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("pvx.xavi_child_sets", param0, param1, param2))
		if "xavi_child_sets" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavi_child_sets']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def xavi_child_rm(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavi_child_rm", param0, param1))
		if "xavi_child_rm" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavi_child_rm']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavi_child_is_null(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.xavi_child_is_null", param0, param1))
		if "xavi_child_is_null" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavi_child_is_null']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def xavi_child_get(self, param0: str, param1: str) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavi_child_get", param0, param1))
		if "xavi_child_get" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavi_child_get']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavi_child_gete(self, param0: str, param1: str) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavi_child_gete", param0, param1))
		if "xavi_child_gete" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavi_child_gete']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def xavi_child_getw(self, param0: str, param1: str) -> Union[int,str]:
		print("Calling %s, %s, %s" % ("pvx.xavi_child_getw", param0, param1))
		if "xavi_child_getw" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['xavi_child_getw']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return None
		return node

	def evalx(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.evalx", param0, param1))
		if "evalx" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['evalx']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def avp_seti(self, param0: str, param1: int) -> int:
		print("Calling %s, %s, %s" % ("pvx.avp_seti", param0, param1))
		if "avp_seti" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['avp_seti']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def avp_sets(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("pvx.avp_sets", param0, param1))
		if "avp_sets" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['avp_sets']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def avp_get(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.avp_get", param0))
		if "avp_get" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['avp_get']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def avp_gete(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.avp_gete", param0))
		if "avp_gete" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['avp_gete']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def avp_getw(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("pvx.avp_getw", param0))
		if "avp_getw" not in _mock_data['pvx']:
			return None
		node = _mock_data['pvx']['avp_getw']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def avp_rm(self, param0: str) -> int:
		print("Calling %s, %s" % ("pvx.avp_rm", param0))
		if "avp_rm" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['avp_rm']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def avp_is_null(self, param0: str) -> int:
		print("Calling %s, %s" % ("pvx.avp_is_null", param0))
		if "avp_is_null" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['avp_is_null']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xavp_copy(self, param0: str, param1: int, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("pvx.xavp_copy", param0, param1, param2))
		if "xavp_copy" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_copy']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def xavp_copy_dst(self, param0: str, param1: int, param2: str, param3: int) -> int:
		print("Calling %s, %s, %s, %s, %s" % ("pvx.xavp_copy_dst", param0, param1, param2, param3))
		if "xavp_copy_dst" not in _mock_data['pvx']:
			return 1
		node = _mock_data['pvx']['xavp_copy_dst']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2, param3)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param3) in node:
			node = node[str(param3)]
		else:
			return 1
		return node

class Maxfwd:
	def process_maxfwd(self, param0: int) -> int:
		print("Calling %s, %s" % ("maxfwd.process_maxfwd", param0))
		if "process_maxfwd" not in _mock_data['maxfwd']:
			return 1
		node = _mock_data['maxfwd']['process_maxfwd']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def is_maxfwd_lt(self, param0: int) -> int:
		print("Calling %s, %s" % ("maxfwd.is_maxfwd_lt", param0))
		if "is_maxfwd_lt" not in _mock_data['maxfwd']:
			return 1
		node = _mock_data['maxfwd']['is_maxfwd_lt']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

class Registrar:
	def save(self, param0: str, param1: int) -> int:
		print("Calling %s, %s, %s" % ("registrar.save", param0, param1))
		if "save" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['save']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def save_uri(self, param0: str, param1: int, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("registrar.save_uri", param0, param1, param2))
		if "save_uri" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['save_uri']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def lookup(self, param0: str) -> int:
		print("Calling %s, %s" % ("registrar.lookup", param0))
		if "lookup" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['lookup']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def lookup_uri(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("registrar.lookup_uri", param0, param1))
		if "lookup_uri" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['lookup_uri']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def lookup_to_dset(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("registrar.lookup_to_dset", param0, param1))
		if "lookup_to_dset" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['lookup_to_dset']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def lookup_branches(self, param0: str) -> int:
		print("Calling %s, %s" % ("registrar.lookup_branches", param0))
		if "lookup_branches" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['lookup_branches']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def registered(self, param0: str) -> int:
		print("Calling %s, %s" % ("registrar.registered", param0))
		if "registered" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['registered']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def registered_uri(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("registrar.registered_uri", param0, param1))
		if "registered_uri" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['registered_uri']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def registered_flags(self, param0: str, param1: str, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("registrar.registered_flags", param0, param1, param2))
		if "registered_flags" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['registered_flags']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def registered_action(self, param0: str, param1: str, param2: int, param3: int) -> int:
		print("Calling %s, %s, %s, %s, %s" % ("registrar.registered_action", param0, param1, param2, param3))
		if "registered_action" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['registered_action']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2, param3)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param3) in node:
			node = node[str(param3)]
		else:
			return 1
		return node

	def set_q_override(self, param0: str) -> int:
		print("Calling %s, %s" % ("registrar.set_q_override", param0))
		if "set_q_override" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['set_q_override']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def add_sock_hdr(self, param0: str) -> int:
		print("Calling %s, %s" % ("registrar.add_sock_hdr", param0))
		if "add_sock_hdr" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['add_sock_hdr']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def unregister(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("registrar.unregister", param0, param1))
		if "unregister" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['unregister']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def unregister_ruid(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("registrar.unregister_ruid", param0, param1, param2))
		if "unregister_ruid" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['unregister_ruid']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def reg_fetch_contacts(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("registrar.reg_fetch_contacts", param0, param1, param2))
		if "reg_fetch_contacts" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['reg_fetch_contacts']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def reg_free_contacts(self, param0: str) -> int:
		print("Calling %s, %s" % ("registrar.reg_free_contacts", param0))
		if "reg_free_contacts" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['reg_free_contacts']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def reg_send_reply(self) -> int:
		print("Calling %s" % ("registrar.reg_send_reply"))
		if "reg_send_reply" not in _mock_data['registrar']:
			return 1
		node = _mock_data['registrar']['reg_send_reply']
		if isinstance(node, types.FunctionType):
			return node()
		return node

class Textops:
	def search(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.search", param0))
		if "search" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['search']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def search_body(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.search_body", param0))
		if "search_body" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['search_body']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def search_hf(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("textops.search_hf", param0, param1, param2))
		if "search_hf" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['search_hf']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def search_append(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.search_append", param0, param1))
		if "search_append" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['search_append']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def search_append_body(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.search_append_body", param0, param1))
		if "search_append_body" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['search_append_body']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def is_present_hf(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.is_present_hf", param0))
		if "is_present_hf" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['is_present_hf']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def is_present_hf_re(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.is_present_hf_re", param0))
		if "is_present_hf_re" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['is_present_hf_re']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def subst(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.subst", param0))
		if "subst" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['subst']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def subst_uri(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.subst_uri", param0))
		if "subst_uri" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['subst_uri']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def subst_user(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.subst_user", param0))
		if "subst_user" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['subst_user']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def subst_body(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.subst_body", param0))
		if "subst_body" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['subst_body']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def subst_hf(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("textops.subst_hf", param0, param1, param2))
		if "subst_hf" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['subst_hf']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def remove_hf(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.remove_hf", param0))
		if "remove_hf" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['remove_hf']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def remove_hf_re(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.remove_hf_re", param0))
		if "remove_hf_re" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['remove_hf_re']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def remove_hf_exp(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.remove_hf_exp", param0, param1))
		if "remove_hf_exp" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['remove_hf_exp']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def replace(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.replace", param0, param1))
		if "replace" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['replace']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def replace_str(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("textops.replace_str", param0, param1, param2))
		if "replace_str" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['replace_str']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def replace_all(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.replace_all", param0, param1))
		if "replace_all" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['replace_all']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def replace_body(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.replace_body", param0, param1))
		if "replace_body" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['replace_body']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def replace_body_str(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("textops.replace_body_str", param0, param1, param2))
		if "replace_body_str" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['replace_body_str']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def replace_hdrs(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.replace_hdrs", param0, param1))
		if "replace_hdrs" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['replace_hdrs']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def replace_hdrs_str(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("textops.replace_hdrs_str", param0, param1, param2))
		if "replace_hdrs_str" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['replace_hdrs_str']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def replace_body_all(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.replace_body_all", param0, param1))
		if "replace_body_all" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['replace_body_all']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def replace_body_atonce(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.replace_body_atonce", param0, param1))
		if "replace_body_atonce" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['replace_body_atonce']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def set_body(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.set_body", param0, param1))
		if "set_body" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['set_body']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def set_reply_body(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.set_reply_body", param0, param1))
		if "set_reply_body" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['set_reply_body']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def has_body(self) -> int:
		print("Calling %s" % ("textops.has_body"))
		if "has_body" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['has_body']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def has_body_type(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.has_body_type", param0))
		if "has_body_type" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['has_body_type']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def filter_body(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.filter_body", param0))
		if "filter_body" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['filter_body']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def is_privacy(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.is_privacy", param0))
		if "is_privacy" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['is_privacy']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def in_list(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("textops.in_list", param0, param1, param2))
		if "in_list" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['in_list']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def in_list_prefix(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("textops.in_list_prefix", param0, param1, param2))
		if "in_list_prefix" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['in_list_prefix']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def cmp_str(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.cmp_str", param0, param1))
		if "cmp_str" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['cmp_str']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def cmp_istr(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.cmp_istr", param0, param1))
		if "cmp_istr" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['cmp_istr']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def search_str(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.search_str", param0, param1))
		if "search_str" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['search_str']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def starts_with(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.starts_with", param0, param1))
		if "starts_with" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['starts_with']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def ends_with(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.ends_with", param0, param1))
		if "ends_with" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['ends_with']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def str_find(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.str_find", param0, param1))
		if "str_find" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['str_find']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def str_ifind(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.str_ifind", param0, param1))
		if "str_ifind" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['str_ifind']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def is_audio_on_hold(self) -> int:
		print("Calling %s" % ("textops.is_audio_on_hold"))
		if "is_audio_on_hold" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['is_audio_on_hold']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def set_body_multipart_mode(self) -> int:
		print("Calling %s" % ("textops.set_body_multipart_mode"))
		if "set_body_multipart_mode" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['set_body_multipart_mode']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def set_body_multipart_boundary(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.set_body_multipart_boundary", param0))
		if "set_body_multipart_boundary" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['set_body_multipart_boundary']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def set_body_multipart_content(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.set_body_multipart_content", param0, param1))
		if "set_body_multipart_content" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['set_body_multipart_content']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def set_body_multipart(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("textops.set_body_multipart", param0, param1, param2))
		if "set_body_multipart" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['set_body_multipart']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def append_body_part(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.append_body_part", param0, param1))
		if "append_body_part" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['append_body_part']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def append_body_part_cd(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("textops.append_body_part_cd", param0, param1, param2))
		if "append_body_part_cd" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['append_body_part_cd']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def append_body_part_hex(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.append_body_part_hex", param0, param1))
		if "append_body_part_hex" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['append_body_part_hex']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def append_body_part_hex_cd(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("textops.append_body_part_hex_cd", param0, param1, param2))
		if "append_body_part_hex_cd" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['append_body_part_hex_cd']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def remove_body_part(self, param0: str) -> int:
		print("Calling %s, %s" % ("textops.remove_body_part", param0))
		if "remove_body_part" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['remove_body_part']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def get_body_part(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.get_body_part", param0, param1))
		if "get_body_part" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['get_body_part']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def get_body_part_raw(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("textops.get_body_part_raw", param0, param1))
		if "get_body_part_raw" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['get_body_part_raw']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def regex_substring(self, param0: str, param1: str, param2: int, param3: int, param4: str) -> int:
		print("Calling %s, %s, %s, %s, %s, %s" % ("textops.regex_substring", param0, param1, param2, param3, param4))
		if "regex_substring" not in _mock_data['textops']:
			return 1
		node = _mock_data['textops']['regex_substring']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2, param3, param4)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param3) in node:
			node = node[str(param3)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param4) in node:
			node = node[str(param4)]
		else:
			return 1
		return node

class Siputils:
	def has_totag(self) -> int:
		print("Calling %s" % ("siputils.has_totag"))
		if "has_totag" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['has_totag']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def is_request(self) -> int:
		print("Calling %s" % ("siputils.is_request"))
		if "is_request" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['is_request']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def is_reply(self) -> int:
		print("Calling %s" % ("siputils.is_reply"))
		if "is_reply" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['is_reply']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def is_first_hop(self) -> int:
		print("Calling %s" % ("siputils.is_first_hop"))
		if "is_first_hop" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['is_first_hop']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def is_uri(self, param0: str) -> int:
		print("Calling %s, %s" % ("siputils.is_uri", param0))
		if "is_uri" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['is_uri']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def is_user(self, param0: str) -> int:
		print("Calling %s, %s" % ("siputils.is_user", param0))
		if "is_user" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['is_user']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def uri_param(self, param0: str) -> int:
		print("Calling %s, %s" % ("siputils.uri_param", param0))
		if "uri_param" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['uri_param']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def uri_param_value(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("siputils.uri_param_value", param0, param1))
		if "uri_param_value" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['uri_param_value']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def uri_param_any(self, param0: str) -> int:
		print("Calling %s, %s" % ("siputils.uri_param_any", param0))
		if "uri_param_any" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['uri_param_any']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def uri_param_rm(self, param0: str) -> int:
		print("Calling %s, %s" % ("siputils.uri_param_rm", param0))
		if "uri_param_rm" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['uri_param_rm']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def is_tel_number(self, param0: str) -> int:
		print("Calling %s, %s" % ("siputils.is_tel_number", param0))
		if "is_tel_number" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['is_tel_number']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def is_numeric(self, param0: str) -> int:
		print("Calling %s, %s" % ("siputils.is_numeric", param0))
		if "is_numeric" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['is_numeric']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def is_alphanum(self, param0: str) -> int:
		print("Calling %s, %s" % ("siputils.is_alphanum", param0))
		if "is_alphanum" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['is_alphanum']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def is_alphanumex(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("siputils.is_alphanumex", param0, param1))
		if "is_alphanumex" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['is_alphanumex']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def options_reply(self) -> int:
		print("Calling %s" % ("siputils.options_reply"))
		if "options_reply" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['options_reply']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def encode_contact(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("siputils.encode_contact", param0, param1))
		if "encode_contact" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['encode_contact']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def decode_contact(self) -> int:
		print("Calling %s" % ("siputils.decode_contact"))
		if "decode_contact" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['decode_contact']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def decode_contact_header(self) -> int:
		print("Calling %s" % ("siputils.decode_contact_header"))
		if "decode_contact_header" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['decode_contact_header']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def contact_param_encode(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("siputils.contact_param_encode", param0, param1))
		if "contact_param_encode" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['contact_param_encode']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def contact_param_decode(self, param0: str) -> int:
		print("Calling %s, %s" % ("siputils.contact_param_decode", param0))
		if "contact_param_decode" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['contact_param_decode']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def contact_param_decode_ruri(self, param0: str) -> int:
		print("Calling %s, %s" % ("siputils.contact_param_decode_ruri", param0))
		if "contact_param_decode_ruri" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['contact_param_decode_ruri']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def contact_param_rm(self, param0: str) -> int:
		print("Calling %s, %s" % ("siputils.contact_param_rm", param0))
		if "contact_param_rm" not in _mock_data['siputils']:
			return 1
		node = _mock_data['siputils']['contact_param_rm']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

class Xlog:
	def xdbg(self, param0: str) -> int:
		print("Calling %s, %s" % ("xlog.xdbg", param0))
		if "xdbg" not in _mock_data['xlog']:
			return 1
		node = _mock_data['xlog']['xdbg']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xinfo(self, param0: str) -> int:
		print("Calling %s, %s" % ("xlog.xinfo", param0))
		if "xinfo" not in _mock_data['xlog']:
			return 1
		node = _mock_data['xlog']['xinfo']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xnotice(self, param0: str) -> int:
		print("Calling %s, %s" % ("xlog.xnotice", param0))
		if "xnotice" not in _mock_data['xlog']:
			return 1
		node = _mock_data['xlog']['xnotice']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xwarn(self, param0: str) -> int:
		print("Calling %s, %s" % ("xlog.xwarn", param0))
		if "xwarn" not in _mock_data['xlog']:
			return 1
		node = _mock_data['xlog']['xwarn']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xerr(self, param0: str) -> int:
		print("Calling %s, %s" % ("xlog.xerr", param0))
		if "xerr" not in _mock_data['xlog']:
			return 1
		node = _mock_data['xlog']['xerr']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xalert(self, param0: str) -> int:
		print("Calling %s, %s" % ("xlog.xalert", param0))
		if "xalert" not in _mock_data['xlog']:
			return 1
		node = _mock_data['xlog']['xalert']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xcrit(self, param0: str) -> int:
		print("Calling %s, %s" % ("xlog.xcrit", param0))
		if "xcrit" not in _mock_data['xlog']:
			return 1
		node = _mock_data['xlog']['xcrit']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def xlog(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("xlog.xlog", param0, param1))
		if "xlog" not in _mock_data['xlog']:
			return 1
		node = _mock_data['xlog']['xlog']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

class Sanity:
	def sanity_check(self, param0: int, param1: int) -> int:
		print("Calling %s, %s, %s" % ("sanity.sanity_check", param0, param1))
		if "sanity_check" not in _mock_data['sanity']:
			return 1
		node = _mock_data['sanity']['sanity_check']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def sanity_check_defaults(self) -> int:
		print("Calling %s" % ("sanity.sanity_check_defaults"))
		if "sanity_check_defaults" not in _mock_data['sanity']:
			return 1
		node = _mock_data['sanity']['sanity_check_defaults']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def sanity_reply(self) -> int:
		print("Calling %s" % ("sanity.sanity_reply"))
		if "sanity_reply" not in _mock_data['sanity']:
			return 1
		node = _mock_data['sanity']['sanity_reply']
		if isinstance(node, types.FunctionType):
			return node()
		return node

class Kex:
	def setdebug(self, param0: int) -> int:
		print("Calling %s, %s" % ("kex.setdebug", param0))
		if "setdebug" not in _mock_data['kex']:
			return 1
		node = _mock_data['kex']['setdebug']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def resetdebug(self) -> int:
		print("Calling %s" % ("kex.resetdebug"))
		if "resetdebug" not in _mock_data['kex']:
			return 1
		node = _mock_data['kex']['resetdebug']
		if isinstance(node, types.FunctionType):
			return node()
		return node

class Corex:
	def append_branch(self) -> int:
		print("Calling %s" % ("corex.append_branch"))
		if "append_branch" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['append_branch']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def append_branch_uri(self, param0: str) -> int:
		print("Calling %s, %s" % ("corex.append_branch_uri", param0))
		if "append_branch_uri" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['append_branch_uri']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def append_branch_uri_q(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("corex.append_branch_uri_q", param0, param1))
		if "append_branch_uri_q" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['append_branch_uri_q']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def setxflag(self, param0: int) -> int:
		print("Calling %s, %s" % ("corex.setxflag", param0))
		if "setxflag" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['setxflag']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def resetxflag(self, param0: int) -> int:
		print("Calling %s, %s" % ("corex.resetxflag", param0))
		if "resetxflag" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['resetxflag']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def isxflagset(self, param0: int) -> int:
		print("Calling %s, %s" % ("corex.isxflagset", param0))
		if "isxflagset" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['isxflagset']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def set_send_socket(self, param0: str) -> int:
		print("Calling %s, %s" % ("corex.set_send_socket", param0))
		if "set_send_socket" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['set_send_socket']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def set_send_socket_name(self, param0: str) -> int:
		print("Calling %s, %s" % ("corex.set_send_socket_name", param0))
		if "set_send_socket_name" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['set_send_socket_name']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def set_recv_socket(self, param0: str) -> int:
		print("Calling %s, %s" % ("corex.set_recv_socket", param0))
		if "set_recv_socket" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['set_recv_socket']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def set_recv_socket_name(self, param0: str) -> int:
		print("Calling %s, %s" % ("corex.set_recv_socket_name", param0))
		if "set_recv_socket_name" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['set_recv_socket_name']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def set_source_address(self, param0: str) -> int:
		print("Calling %s, %s" % ("corex.set_source_address", param0))
		if "set_source_address" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['set_source_address']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def via_add_srvid(self, param0: int) -> int:
		print("Calling %s, %s" % ("corex.via_add_srvid", param0))
		if "via_add_srvid" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['via_add_srvid']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def via_add_xavp_params(self, param0: int) -> int:
		print("Calling %s, %s" % ("corex.via_add_xavp_params", param0))
		if "via_add_xavp_params" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['via_add_xavp_params']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def via_use_xavp_fields(self, param0: int) -> int:
		print("Calling %s, %s" % ("corex.via_use_xavp_fields", param0))
		if "via_use_xavp_fields" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['via_use_xavp_fields']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def has_ruri_user(self) -> int:
		print("Calling %s" % ("corex.has_ruri_user"))
		if "has_ruri_user" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['has_ruri_user']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def has_user_agent(self) -> int:
		print("Calling %s" % ("corex.has_user_agent"))
		if "has_user_agent" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['has_user_agent']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def send_data(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("corex.send_data", param0, param1))
		if "send_data" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['send_data']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def sendx(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("corex.sendx", param0, param1, param2))
		if "sendx" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['sendx']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def is_faked_msg(self) -> int:
		print("Calling %s" % ("corex.is_faked_msg"))
		if "is_faked_msg" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['is_faked_msg']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def file_read(self, param0: str) -> Union[int,str]:
		print("Calling %s, %s" % ("corex.file_read", param0))
		if "file_read" not in _mock_data['corex']:
			return None
		node = _mock_data['corex']['file_read']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return None
		return node

	def file_write(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("corex.file_write", param0, param1))
		if "file_write" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['file_write']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def is_socket_name(self, param0: str) -> int:
		print("Calling %s, %s" % ("corex.is_socket_name", param0))
		if "is_socket_name" not in _mock_data['corex']:
			return 1
		node = _mock_data['corex']['is_socket_name']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

class Jsonrpcs:
	def dispatch(self) -> int:
		print("Calling %s" % ("jsonrpcs.dispatch"))
		if "dispatch" not in _mock_data['jsonrpcs']:
			return 1
		node = _mock_data['jsonrpcs']['dispatch']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def exec(self, param0: str) -> int:
		print("Calling %s, %s" % ("jsonrpcs.exec", param0))
		if "exec" not in _mock_data['jsonrpcs']:
			return 1
		node = _mock_data['jsonrpcs']['exec']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def execx(self, param0: str) -> int:
		print("Calling %s, %s" % ("jsonrpcs.execx", param0))
		if "execx" not in _mock_data['jsonrpcs']:
			return 1
		node = _mock_data['jsonrpcs']['execx']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def response(self) -> Union[int,str]:
		print("Calling %s" % ("jsonrpcs.response"))
		if "response" not in _mock_data['jsonrpcs']:
			return None
		node = _mock_data['jsonrpcs']['response']
		if isinstance(node, types.FunctionType):
			return node()
		return node

class Http_async_client:
	def query(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("http_async_client.query", param0, param1))
		if "query" not in _mock_data['http_async_client']:
			return 1
		node = _mock_data['http_async_client']['query']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

class Xhttp:
	def xhttp_reply(self, param0: int, param1: str, param2: str, param3: str) -> int:
		print("Calling %s, %s, %s, %s, %s" % ("xhttp.xhttp_reply", param0, param1, param2, param3))
		if "xhttp_reply" not in _mock_data['xhttp']:
			return 1
		node = _mock_data['xhttp']['xhttp_reply']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2, param3)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param3) in node:
			node = node[str(param3)]
		else:
			return 1
		return node

	def get_hu(self) -> Union[int,str]:
		print("Calling %s" % ("xhttp.get_hu"))
		if "get_hu" not in _mock_data['xhttp']:
			return None
		node = _mock_data['xhttp']['get_hu']
		if isinstance(node, types.FunctionType):
			return node()
		return node

class Websocket:
	def handle_handshake(self) -> int:
		print("Calling %s" % ("websocket.handle_handshake"))
		if "handle_handshake" not in _mock_data['websocket']:
			return 1
		node = _mock_data['websocket']['handle_handshake']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def close(self) -> int:
		print("Calling %s" % ("websocket.close"))
		if "close" not in _mock_data['websocket']:
			return 1
		node = _mock_data['websocket']['close']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def close_reason(self, param0: int, param1: str) -> int:
		print("Calling %s, %s, %s" % ("websocket.close_reason", param0, param1))
		if "close_reason" not in _mock_data['websocket']:
			return 1
		node = _mock_data['websocket']['close_reason']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def close_conid(self, param0: int, param1: str, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("websocket.close_conid", param0, param1, param2))
		if "close_conid" not in _mock_data['websocket']:
			return 1
		node = _mock_data['websocket']['close_conid']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

class Nathelper:
	def nat_uac_test(self, param0: int) -> int:
		print("Calling %s, %s" % ("nathelper.nat_uac_test", param0))
		if "nat_uac_test" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['nat_uac_test']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def fix_nated_contact(self) -> int:
		print("Calling %s" % ("nathelper.fix_nated_contact"))
		if "fix_nated_contact" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['fix_nated_contact']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def fix_nated_register(self) -> int:
		print("Calling %s" % ("nathelper.fix_nated_register"))
		if "fix_nated_register" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['fix_nated_register']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def set_contact_alias(self) -> int:
		print("Calling %s" % ("nathelper.set_contact_alias"))
		if "set_contact_alias" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['set_contact_alias']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def set_contact_alias_trim(self) -> int:
		print("Calling %s" % ("nathelper.set_contact_alias_trim"))
		if "set_contact_alias_trim" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['set_contact_alias_trim']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def handle_ruri_alias(self) -> int:
		print("Calling %s" % ("nathelper.handle_ruri_alias"))
		if "handle_ruri_alias" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['handle_ruri_alias']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def is_rfc1918(self, param0: str) -> int:
		print("Calling %s, %s" % ("nathelper.is_rfc1918", param0))
		if "is_rfc1918" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['is_rfc1918']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def add_contact_alias(self) -> int:
		print("Calling %s" % ("nathelper.add_contact_alias"))
		if "add_contact_alias" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['add_contact_alias']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def add_contact_alias_addr(self, param0: str, param1: str, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("nathelper.add_contact_alias_addr", param0, param1, param2))
		if "add_contact_alias_addr" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['add_contact_alias_addr']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def add_rcv_param(self, param0: int) -> int:
		print("Calling %s, %s" % ("nathelper.add_rcv_param", param0))
		if "add_rcv_param" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['add_rcv_param']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def fix_nated_sdp(self, param0: int) -> int:
		print("Calling %s, %s" % ("nathelper.fix_nated_sdp", param0))
		if "fix_nated_sdp" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['fix_nated_sdp']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def fix_nated_sdp_ip(self, param0: int, param1: str) -> int:
		print("Calling %s, %s, %s" % ("nathelper.fix_nated_sdp_ip", param0, param1))
		if "fix_nated_sdp_ip" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['fix_nated_sdp_ip']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def set_alias_to_pv(self, param0: str) -> int:
		print("Calling %s, %s" % ("nathelper.set_alias_to_pv", param0))
		if "set_alias_to_pv" not in _mock_data['nathelper']:
			return 1
		node = _mock_data['nathelper']['set_alias_to_pv']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

class App_python3:
	def exec(self, param0: str) -> int:
		print("Calling %s, %s" % ("app_python3.exec", param0))
		if "exec" not in _mock_data['app_python3']:
			return 1
		node = _mock_data['app_python3']['exec']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def execx(self, param0: str) -> int:
		print("Calling %s, %s" % ("app_python3.execx", param0))
		if "execx" not in _mock_data['app_python3']:
			return 1
		node = _mock_data['app_python3']['execx']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def exec_p1(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("app_python3.exec_p1", param0, param1))
		if "exec_p1" not in _mock_data['app_python3']:
			return 1
		node = _mock_data['app_python3']['exec_p1']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

class Auth:
	def consume_credentials(self) -> int:
		print("Calling %s" % ("auth.consume_credentials"))
		if "consume_credentials" not in _mock_data['auth']:
			return 1
		node = _mock_data['auth']['consume_credentials']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def auth_challenge(self, param0: str, param1: int) -> int:
		print("Calling %s, %s, %s" % ("auth.auth_challenge", param0, param1))
		if "auth_challenge" not in _mock_data['auth']:
			return 1
		node = _mock_data['auth']['auth_challenge']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def pv_auth_check(self, param0: str, param1: str, param2: int, param3: int) -> int:
		print("Calling %s, %s, %s, %s, %s" % ("auth.pv_auth_check", param0, param1, param2, param3))
		if "pv_auth_check" not in _mock_data['auth']:
			return 1
		node = _mock_data['auth']['pv_auth_check']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2, param3)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param3) in node:
			node = node[str(param3)]
		else:
			return 1
		return node

	def has_credentials(self, param0: str) -> int:
		print("Calling %s, %s" % ("auth.has_credentials", param0))
		if "has_credentials" not in _mock_data['auth']:
			return 1
		node = _mock_data['auth']['has_credentials']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

class Auth_ephemeral:
	def autheph_check(self, param0: str) -> int:
		print("Calling %s, %s" % ("auth_ephemeral.autheph_check", param0))
		if "autheph_check" not in _mock_data['auth_ephemeral']:
			return 1
		node = _mock_data['auth_ephemeral']['autheph_check']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def autheph_www(self, param0: str) -> int:
		print("Calling %s, %s" % ("auth_ephemeral.autheph_www", param0))
		if "autheph_www" not in _mock_data['auth_ephemeral']:
			return 1
		node = _mock_data['auth_ephemeral']['autheph_www']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def autheph_www_method(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("auth_ephemeral.autheph_www_method", param0, param1))
		if "autheph_www_method" not in _mock_data['auth_ephemeral']:
			return 1
		node = _mock_data['auth_ephemeral']['autheph_www_method']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def autheph_proxy(self, param0: str) -> int:
		print("Calling %s, %s" % ("auth_ephemeral.autheph_proxy", param0))
		if "autheph_proxy" not in _mock_data['auth_ephemeral']:
			return 1
		node = _mock_data['auth_ephemeral']['autheph_proxy']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def autheph_authenticate(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("auth_ephemeral.autheph_authenticate", param0, param1))
		if "autheph_authenticate" not in _mock_data['auth_ephemeral']:
			return 1
		node = _mock_data['auth_ephemeral']['autheph_authenticate']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

class Dispatcher:
	def ds_select(self, param0: int, param1: int) -> int:
		print("Calling %s, %s, %s" % ("dispatcher.ds_select", param0, param1))
		if "ds_select" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_select']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def ds_select_limit(self, param0: int, param1: int, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("dispatcher.ds_select_limit", param0, param1, param2))
		if "ds_select_limit" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_select_limit']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def ds_select_domain(self, param0: int, param1: int) -> int:
		print("Calling %s, %s, %s" % ("dispatcher.ds_select_domain", param0, param1))
		if "ds_select_domain" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_select_domain']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def ds_select_domain_limit(self, param0: int, param1: int, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("dispatcher.ds_select_domain_limit", param0, param1, param2))
		if "ds_select_domain_limit" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_select_domain_limit']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def ds_next_domain(self) -> int:
		print("Calling %s" % ("dispatcher.ds_next_domain"))
		if "ds_next_domain" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_next_domain']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def ds_set_domain(self) -> int:
		print("Calling %s" % ("dispatcher.ds_set_domain"))
		if "ds_set_domain" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_set_domain']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def ds_select_dst(self, param0: int, param1: int) -> int:
		print("Calling %s, %s, %s" % ("dispatcher.ds_select_dst", param0, param1))
		if "ds_select_dst" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_select_dst']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def ds_select_dst_limit(self, param0: int, param1: int, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("dispatcher.ds_select_dst_limit", param0, param1, param2))
		if "ds_select_dst_limit" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_select_dst_limit']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def ds_select_routes(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("dispatcher.ds_select_routes", param0, param1))
		if "ds_select_routes" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_select_routes']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def ds_select_routes_limit(self, param0: str, param1: str, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("dispatcher.ds_select_routes_limit", param0, param1, param2))
		if "ds_select_routes_limit" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_select_routes_limit']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def ds_next_dst(self) -> int:
		print("Calling %s" % ("dispatcher.ds_next_dst"))
		if "ds_next_dst" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_next_dst']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def ds_set_dst(self) -> int:
		print("Calling %s" % ("dispatcher.ds_set_dst"))
		if "ds_set_dst" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_set_dst']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def ds_mark_dst(self) -> int:
		print("Calling %s" % ("dispatcher.ds_mark_dst"))
		if "ds_mark_dst" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_mark_dst']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def ds_mark_dst_state(self, param0: str) -> int:
		print("Calling %s, %s" % ("dispatcher.ds_mark_dst_state", param0))
		if "ds_mark_dst_state" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_mark_dst_state']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def ds_is_from_lists(self) -> int:
		print("Calling %s" % ("dispatcher.ds_is_from_lists"))
		if "ds_is_from_lists" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_is_from_lists']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def ds_is_from_list(self, param0: int) -> int:
		print("Calling %s, %s" % ("dispatcher.ds_is_from_list", param0))
		if "ds_is_from_list" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_is_from_list']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def ds_is_from_list_mode(self, param0: int, param1: int) -> int:
		print("Calling %s, %s, %s" % ("dispatcher.ds_is_from_list_mode", param0, param1))
		if "ds_is_from_list_mode" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_is_from_list_mode']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def ds_is_from_list_uri(self, param0: int, param1: int, param2: str) -> int:
		print("Calling %s, %s, %s, %s" % ("dispatcher.ds_is_from_list_uri", param0, param1, param2))
		if "ds_is_from_list_uri" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_is_from_list_uri']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def ds_load_update(self) -> int:
		print("Calling %s" % ("dispatcher.ds_load_update"))
		if "ds_load_update" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_load_update']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def ds_load_unset(self) -> int:
		print("Calling %s" % ("dispatcher.ds_load_unset"))
		if "ds_load_unset" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_load_unset']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def ds_reload(self) -> int:
		print("Calling %s" % ("dispatcher.ds_reload"))
		if "ds_reload" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_reload']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def ds_list_exists(self, param0: int) -> int:
		print("Calling %s, %s" % ("dispatcher.ds_list_exists", param0))
		if "ds_list_exists" not in _mock_data['dispatcher']:
			return 1
		node = _mock_data['dispatcher']['ds_list_exists']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

class Uac:
	def uac_auth(self) -> int:
		print("Calling %s" % ("uac.uac_auth"))
		if "uac_auth" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_auth']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def uac_auth_mode(self, param0: int) -> int:
		print("Calling %s, %s" % ("uac.uac_auth_mode", param0))
		if "uac_auth_mode" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_auth_mode']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def uac_req_send(self) -> int:
		print("Calling %s" % ("uac.uac_req_send"))
		if "uac_req_send" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_req_send']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def uac_replace_from_uri(self, param0: str) -> int:
		print("Calling %s, %s" % ("uac.uac_replace_from_uri", param0))
		if "uac_replace_from_uri" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_replace_from_uri']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def uac_replace_from(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("uac.uac_replace_from", param0, param1))
		if "uac_replace_from" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_replace_from']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def uac_restore_from(self) -> int:
		print("Calling %s" % ("uac.uac_restore_from"))
		if "uac_restore_from" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_restore_from']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def uac_replace_to_uri(self, param0: str) -> int:
		print("Calling %s, %s" % ("uac.uac_replace_to_uri", param0))
		if "uac_replace_to_uri" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_replace_to_uri']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def uac_replace_to(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("uac.uac_replace_to", param0, param1))
		if "uac_replace_to" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_replace_to']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def uac_restore_to(self) -> int:
		print("Calling %s" % ("uac.uac_restore_to"))
		if "uac_restore_to" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_restore_to']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def uac_reg_lookup(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("uac.uac_reg_lookup", param0, param1))
		if "uac_reg_lookup" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_reg_lookup']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def uac_reg_status(self, param0: str) -> int:
		print("Calling %s, %s" % ("uac.uac_reg_status", param0))
		if "uac_reg_status" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_reg_status']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def uac_reg_request_to(self, param0: str, param1: int) -> int:
		print("Calling %s, %s, %s" % ("uac.uac_reg_request_to", param0, param1))
		if "uac_reg_request_to" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_reg_request_to']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def uac_reg_enable(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("uac.uac_reg_enable", param0, param1))
		if "uac_reg_enable" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_reg_enable']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def uac_reg_disable(self, param0: str, param1: str) -> int:
		print("Calling %s, %s, %s" % ("uac.uac_reg_disable", param0, param1))
		if "uac_reg_disable" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_reg_disable']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def uac_reg_refresh(self, param0: str) -> int:
		print("Calling %s, %s" % ("uac.uac_reg_refresh", param0))
		if "uac_reg_refresh" not in _mock_data['uac']:
			return 1
		node = _mock_data['uac']['uac_reg_refresh']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

class Permissions:
	def allow_source_address(self, param0: int) -> int:
		print("Calling %s, %s" % ("permissions.allow_source_address", param0))
		if "allow_source_address" not in _mock_data['permissions']:
			return 1
		node = _mock_data['permissions']['allow_source_address']
		if isinstance(node, types.FunctionType):
			return node(param0)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		return node

	def allow_address(self, param0: int, param1: str, param2: int) -> int:
		print("Calling %s, %s, %s, %s" % ("permissions.allow_address", param0, param1, param2))
		if "allow_address" not in _mock_data['permissions']:
			return 1
		node = _mock_data['permissions']['allow_address']
		if isinstance(node, types.FunctionType):
			return node(param0, param1, param2)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param2) in node:
			node = node[str(param2)]
		else:
			return 1
		return node

	def allow_source_address_group(self) -> int:
		print("Calling %s" % ("permissions.allow_source_address_group"))
		if "allow_source_address_group" not in _mock_data['permissions']:
			return 1
		node = _mock_data['permissions']['allow_source_address_group']
		if isinstance(node, types.FunctionType):
			return node()
		return node

	def allow_address_group(self, param0: str, param1: int) -> int:
		print("Calling %s, %s, %s" % ("permissions.allow_address_group", param0, param1))
		if "allow_address_group" not in _mock_data['permissions']:
			return 1
		node = _mock_data['permissions']['allow_address_group']
		if isinstance(node, types.FunctionType):
			return node(param0, param1)
		if not isinstance(node, dict):
			return node
		if str(param0) in node:
			node = node[str(param0)]
		else:
			return 1
		if not isinstance(node, dict):
			return node
		if str(param1) in node:
			node = node[str(param1)]
		else:
			return 1
		return node

	def allow_trusted(self) -> int:
		print("Calling %s" % ("permissions.allow_trusted"))
		if "allow_trusted" not in _mock_data['permissions']:
			return 1
		node = _mock_data['permissions']['allow_trusted']
		if isinstance(node, types.FunctionType):
			return node()
		return node


def dbg(param0: str):
	print("Calling %s, %s" % ("dbg", param0))
	if "dbg" not in _mock_data['']:
		return
	node = _mock_data['']['dbg']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return
	return node


def err(param0: str):
	print("Calling %s, %s" % ("err", param0))
	if "err" not in _mock_data['']:
		return
	node = _mock_data['']['err']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return
	return node


def info(param0: str):
	print("Calling %s, %s" % ("info", param0))
	if "info" not in _mock_data['']:
		return
	node = _mock_data['']['info']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return
	return node


def warn(param0: str):
	print("Calling %s, %s" % ("warn", param0))
	if "warn" not in _mock_data['']:
		return
	node = _mock_data['']['warn']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return
	return node


def notice(param0: str):
	print("Calling %s, %s" % ("notice", param0))
	if "notice" not in _mock_data['']:
		return
	node = _mock_data['']['notice']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return
	return node


def crit(param0: str):
	print("Calling %s, %s" % ("crit", param0))
	if "crit" not in _mock_data['']:
		return
	node = _mock_data['']['crit']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return
	return node


def log(param0: str, param1: str):
	print("Calling %s, %s, %s" % ("log", param0, param1))
	if "log" not in _mock_data['']:
		return
	node = _mock_data['']['log']
	if isinstance(node, types.FunctionType):
		return node(param0, param1)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return
	if not isinstance(node, dict):
		return node
	if str(param1) in node:
		node = node[str(param1)]
	else:
		return
	return node


def set_drop():
	print("Calling %s" % ("set_drop"))
	if "set_drop" not in _mock_data['']:
		return
	node = _mock_data['']['set_drop']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_myself(param0: str) -> bool:
	print("Calling %s, %s" % ("is_myself", param0))
	if "is_myself" not in _mock_data['']:
		return True
	node = _mock_data['']['is_myself']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def is_myself_ruri() -> bool:
	print("Calling %s" % ("is_myself_ruri"))
	if "is_myself_ruri" not in _mock_data['']:
		return True
	node = _mock_data['']['is_myself_ruri']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_myself_duri() -> bool:
	print("Calling %s" % ("is_myself_duri"))
	if "is_myself_duri" not in _mock_data['']:
		return True
	node = _mock_data['']['is_myself_duri']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_myself_nhuri() -> bool:
	print("Calling %s" % ("is_myself_nhuri"))
	if "is_myself_nhuri" not in _mock_data['']:
		return True
	node = _mock_data['']['is_myself_nhuri']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_myself_furi() -> bool:
	print("Calling %s" % ("is_myself_furi"))
	if "is_myself_furi" not in _mock_data['']:
		return True
	node = _mock_data['']['is_myself_furi']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_myself_turi() -> bool:
	print("Calling %s" % ("is_myself_turi"))
	if "is_myself_turi" not in _mock_data['']:
		return True
	node = _mock_data['']['is_myself_turi']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_myself_suri() -> bool:
	print("Calling %s" % ("is_myself_suri"))
	if "is_myself_suri" not in _mock_data['']:
		return True
	node = _mock_data['']['is_myself_suri']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_myself_srcip() -> bool:
	print("Calling %s" % ("is_myself_srcip"))
	if "is_myself_srcip" not in _mock_data['']:
		return True
	node = _mock_data['']['is_myself_srcip']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def setflag(param0: int) -> bool:
	print("Calling %s, %s" % ("setflag", param0))
	if "setflag" not in _mock_data['']:
		return True
	node = _mock_data['']['setflag']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def resetflag(param0: int) -> bool:
	print("Calling %s, %s" % ("resetflag", param0))
	if "resetflag" not in _mock_data['']:
		return True
	node = _mock_data['']['resetflag']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def isflagset(param0: int) -> bool:
	print("Calling %s, %s" % ("isflagset", param0))
	if "isflagset" not in _mock_data['']:
		return True
	node = _mock_data['']['isflagset']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def setbflag(param0: int) -> bool:
	print("Calling %s, %s" % ("setbflag", param0))
	if "setbflag" not in _mock_data['']:
		return True
	node = _mock_data['']['setbflag']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def resetbflag(param0: int) -> bool:
	print("Calling %s, %s" % ("resetbflag", param0))
	if "resetbflag" not in _mock_data['']:
		return True
	node = _mock_data['']['resetbflag']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def isbflagset(param0: int) -> bool:
	print("Calling %s, %s" % ("isbflagset", param0))
	if "isbflagset" not in _mock_data['']:
		return True
	node = _mock_data['']['isbflagset']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def setbiflag(param0: int, param1: int) -> bool:
	print("Calling %s, %s, %s" % ("setbiflag", param0, param1))
	if "setbiflag" not in _mock_data['']:
		return True
	node = _mock_data['']['setbiflag']
	if isinstance(node, types.FunctionType):
		return node(param0, param1)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	if not isinstance(node, dict):
		return node
	if str(param1) in node:
		node = node[str(param1)]
	else:
		return True
	return node


def resetbiflag(param0: int, param1: int) -> bool:
	print("Calling %s, %s, %s" % ("resetbiflag", param0, param1))
	if "resetbiflag" not in _mock_data['']:
		return True
	node = _mock_data['']['resetbiflag']
	if isinstance(node, types.FunctionType):
		return node(param0, param1)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	if not isinstance(node, dict):
		return node
	if str(param1) in node:
		node = node[str(param1)]
	else:
		return True
	return node


def isbiflagset(param0: int, param1: int) -> bool:
	print("Calling %s, %s, %s" % ("isbiflagset", param0, param1))
	if "isbiflagset" not in _mock_data['']:
		return True
	node = _mock_data['']['isbiflagset']
	if isinstance(node, types.FunctionType):
		return node(param0, param1)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	if not isinstance(node, dict):
		return node
	if str(param1) in node:
		node = node[str(param1)]
	else:
		return True
	return node


def setsflag(param0: int) -> bool:
	print("Calling %s, %s" % ("setsflag", param0))
	if "setsflag" not in _mock_data['']:
		return True
	node = _mock_data['']['setsflag']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def resetsflag(param0: int) -> bool:
	print("Calling %s, %s" % ("resetsflag", param0))
	if "resetsflag" not in _mock_data['']:
		return True
	node = _mock_data['']['resetsflag']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def issflagset(param0: int) -> bool:
	print("Calling %s, %s" % ("issflagset", param0))
	if "issflagset" not in _mock_data['']:
		return True
	node = _mock_data['']['issflagset']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def seturi(param0: str) -> bool:
	print("Calling %s, %s" % ("seturi", param0))
	if "seturi" not in _mock_data['']:
		return True
	node = _mock_data['']['seturi']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def setuser(param0: str) -> bool:
	print("Calling %s, %s" % ("setuser", param0))
	if "setuser" not in _mock_data['']:
		return True
	node = _mock_data['']['setuser']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def sethost(param0: str) -> bool:
	print("Calling %s, %s" % ("sethost", param0))
	if "sethost" not in _mock_data['']:
		return True
	node = _mock_data['']['sethost']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def setdsturi(param0: str) -> bool:
	print("Calling %s, %s" % ("setdsturi", param0))
	if "setdsturi" not in _mock_data['']:
		return True
	node = _mock_data['']['setdsturi']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def resetdsturi() -> bool:
	print("Calling %s" % ("resetdsturi"))
	if "resetdsturi" not in _mock_data['']:
		return True
	node = _mock_data['']['resetdsturi']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def isdsturiset() -> bool:
	print("Calling %s" % ("isdsturiset"))
	if "isdsturiset" not in _mock_data['']:
		return True
	node = _mock_data['']['isdsturiset']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def force_rport() -> bool:
	print("Calling %s" % ("force_rport"))
	if "force_rport" not in _mock_data['']:
		return True
	node = _mock_data['']['force_rport']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def add_local_rport() -> bool:
	print("Calling %s" % ("add_local_rport"))
	if "add_local_rport" not in _mock_data['']:
		return True
	node = _mock_data['']['add_local_rport']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_method(param0: str) -> bool:
	print("Calling %s, %s" % ("is_method", param0))
	if "is_method" not in _mock_data['']:
		return True
	node = _mock_data['']['is_method']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def is_method_in(param0: str) -> bool:
	print("Calling %s, %s" % ("is_method_in", param0))
	if "is_method_in" not in _mock_data['']:
		return True
	node = _mock_data['']['is_method_in']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def forward() -> int:
	print("Calling %s" % ("forward"))
	if "forward" not in _mock_data['']:
		return 1
	node = _mock_data['']['forward']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def forward_uri(param0: str) -> int:
	print("Calling %s, %s" % ("forward_uri", param0))
	if "forward_uri" not in _mock_data['']:
		return 1
	node = _mock_data['']['forward_uri']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return 1
	return node


def set_forward_close() -> bool:
	print("Calling %s" % ("set_forward_close"))
	if "set_forward_close" not in _mock_data['']:
		return True
	node = _mock_data['']['set_forward_close']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def set_forward_no_connect() -> bool:
	print("Calling %s" % ("set_forward_no_connect"))
	if "set_forward_no_connect" not in _mock_data['']:
		return True
	node = _mock_data['']['set_forward_no_connect']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def set_reply_close() -> bool:
	print("Calling %s" % ("set_reply_close"))
	if "set_reply_close" not in _mock_data['']:
		return True
	node = _mock_data['']['set_reply_close']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def set_reply_no_connect() -> bool:
	print("Calling %s" % ("set_reply_no_connect"))
	if "set_reply_no_connect" not in _mock_data['']:
		return True
	node = _mock_data['']['set_reply_no_connect']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def set_advertised_address(param0: str) -> int:
	print("Calling %s, %s" % ("set_advertised_address", param0))
	if "set_advertised_address" not in _mock_data['']:
		return 1
	node = _mock_data['']['set_advertised_address']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return 1
	return node


def set_advertised_port(param0: str) -> int:
	print("Calling %s, %s" % ("set_advertised_port", param0))
	if "set_advertised_port" not in _mock_data['']:
		return 1
	node = _mock_data['']['set_advertised_port']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return 1
	return node


def add_tcp_alias(param0: int) -> int:
	print("Calling %s, %s" % ("add_tcp_alias", param0))
	if "add_tcp_alias" not in _mock_data['']:
		return 1
	node = _mock_data['']['add_tcp_alias']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return 1
	return node


def add_tcp_alias_via() -> int:
	print("Calling %s" % ("add_tcp_alias_via"))
	if "add_tcp_alias_via" not in _mock_data['']:
		return 1
	node = _mock_data['']['add_tcp_alias_via']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_INVITE() -> bool:
	print("Calling %s" % ("is_INVITE"))
	if "is_INVITE" not in _mock_data['']:
		return True
	node = _mock_data['']['is_INVITE']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_ACK() -> bool:
	print("Calling %s" % ("is_ACK"))
	if "is_ACK" not in _mock_data['']:
		return True
	node = _mock_data['']['is_ACK']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_BYE() -> bool:
	print("Calling %s" % ("is_BYE"))
	if "is_BYE" not in _mock_data['']:
		return True
	node = _mock_data['']['is_BYE']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_CANCEL() -> bool:
	print("Calling %s" % ("is_CANCEL"))
	if "is_CANCEL" not in _mock_data['']:
		return True
	node = _mock_data['']['is_CANCEL']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_REGISTER() -> bool:
	print("Calling %s" % ("is_REGISTER"))
	if "is_REGISTER" not in _mock_data['']:
		return True
	node = _mock_data['']['is_REGISTER']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_OPTIONS() -> bool:
	print("Calling %s" % ("is_OPTIONS"))
	if "is_OPTIONS" not in _mock_data['']:
		return True
	node = _mock_data['']['is_OPTIONS']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_SUBSCRIBE() -> bool:
	print("Calling %s" % ("is_SUBSCRIBE"))
	if "is_SUBSCRIBE" not in _mock_data['']:
		return True
	node = _mock_data['']['is_SUBSCRIBE']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_PUBLISH() -> bool:
	print("Calling %s" % ("is_PUBLISH"))
	if "is_PUBLISH" not in _mock_data['']:
		return True
	node = _mock_data['']['is_PUBLISH']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_NOTIFY() -> bool:
	print("Calling %s" % ("is_NOTIFY"))
	if "is_NOTIFY" not in _mock_data['']:
		return True
	node = _mock_data['']['is_NOTIFY']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_REFER() -> bool:
	print("Calling %s" % ("is_REFER"))
	if "is_REFER" not in _mock_data['']:
		return True
	node = _mock_data['']['is_REFER']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_INFO() -> bool:
	print("Calling %s" % ("is_INFO"))
	if "is_INFO" not in _mock_data['']:
		return True
	node = _mock_data['']['is_INFO']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_UPDATE() -> bool:
	print("Calling %s" % ("is_UPDATE"))
	if "is_UPDATE" not in _mock_data['']:
		return True
	node = _mock_data['']['is_UPDATE']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_PRACK() -> bool:
	print("Calling %s" % ("is_PRACK"))
	if "is_PRACK" not in _mock_data['']:
		return True
	node = _mock_data['']['is_PRACK']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_MESSAGE() -> bool:
	print("Calling %s" % ("is_MESSAGE"))
	if "is_MESSAGE" not in _mock_data['']:
		return True
	node = _mock_data['']['is_MESSAGE']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_KDMQ() -> bool:
	print("Calling %s" % ("is_KDMQ"))
	if "is_KDMQ" not in _mock_data['']:
		return True
	node = _mock_data['']['is_KDMQ']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_GET() -> bool:
	print("Calling %s" % ("is_GET"))
	if "is_GET" not in _mock_data['']:
		return True
	node = _mock_data['']['is_GET']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_POST() -> bool:
	print("Calling %s" % ("is_POST"))
	if "is_POST" not in _mock_data['']:
		return True
	node = _mock_data['']['is_POST']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_PUT() -> bool:
	print("Calling %s" % ("is_PUT"))
	if "is_PUT" not in _mock_data['']:
		return True
	node = _mock_data['']['is_PUT']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_DELETE() -> bool:
	print("Calling %s" % ("is_DELETE"))
	if "is_DELETE" not in _mock_data['']:
		return True
	node = _mock_data['']['is_DELETE']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_UDP() -> bool:
	print("Calling %s" % ("is_UDP"))
	if "is_UDP" not in _mock_data['']:
		return True
	node = _mock_data['']['is_UDP']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_TCP() -> bool:
	print("Calling %s" % ("is_TCP"))
	if "is_TCP" not in _mock_data['']:
		return True
	node = _mock_data['']['is_TCP']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_TLS() -> bool:
	print("Calling %s" % ("is_TLS"))
	if "is_TLS" not in _mock_data['']:
		return True
	node = _mock_data['']['is_TLS']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_WS() -> bool:
	print("Calling %s" % ("is_WS"))
	if "is_WS" not in _mock_data['']:
		return True
	node = _mock_data['']['is_WS']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_WSS() -> bool:
	print("Calling %s" % ("is_WSS"))
	if "is_WSS" not in _mock_data['']:
		return True
	node = _mock_data['']['is_WSS']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_SCTP() -> bool:
	print("Calling %s" % ("is_SCTP"))
	if "is_SCTP" not in _mock_data['']:
		return True
	node = _mock_data['']['is_SCTP']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_proto(param0: str) -> bool:
	print("Calling %s, %s" % ("is_proto", param0))
	if "is_proto" not in _mock_data['']:
		return True
	node = _mock_data['']['is_proto']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def is_IPv4() -> bool:
	print("Calling %s" % ("is_IPv4"))
	if "is_IPv4" not in _mock_data['']:
		return True
	node = _mock_data['']['is_IPv4']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_IPv6() -> bool:
	print("Calling %s" % ("is_IPv6"))
	if "is_IPv6" not in _mock_data['']:
		return True
	node = _mock_data['']['is_IPv6']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def is_src_port(param0: int) -> bool:
	print("Calling %s, %s" % ("is_src_port", param0))
	if "is_src_port" not in _mock_data['']:
		return True
	node = _mock_data['']['is_src_port']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def is_dst_port(param0: int) -> bool:
	print("Calling %s, %s" % ("is_dst_port", param0))
	if "is_dst_port" not in _mock_data['']:
		return True
	node = _mock_data['']['is_dst_port']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return True
	return node


def get_debug() -> int:
	print("Calling %s" % ("get_debug"))
	if "get_debug" not in _mock_data['']:
		return 1
	node = _mock_data['']['get_debug']
	if isinstance(node, types.FunctionType):
		return node()
	return node


def route(param0: str) -> int:
	print("Calling %s, %s" % ("route", param0))
	if "route" not in _mock_data['']:
		return 1
	node = _mock_data['']['route']
	if isinstance(node, types.FunctionType):
		return node(param0)
	if not isinstance(node, dict):
		return node
	if str(param0) in node:
		node = node[str(param0)]
	else:
		return 1
	return node

pv = Pv()
hdr = Hdr()
tm = Tm()
tmx = Tmx()
sl = Sl()
rr = Rr()
pvx = Pvx()
maxfwd = Maxfwd()
registrar = Registrar()
textops = Textops()
siputils = Siputils()
xlog = Xlog()
sanity = Sanity()
kex = Kex()
corex = Corex()
jsonrpcs = Jsonrpcs()
http_async_client = Http_async_client()
xhttp = Xhttp()
websocket = Websocket()
nathelper = Nathelper()
app_python3 = App_python3()
auth = Auth()
auth_ephemeral = Auth_ephemeral()
dispatcher = Dispatcher()
uac = Uac()
permissions = Permissions()

_mock_data[''] = {}
_mock_data['pv'] = {}
_mock_data['hdr'] = {}
_mock_data['tm'] = {}
_mock_data['tmx'] = {}
_mock_data['sl'] = {}
_mock_data['rr'] = {}
_mock_data['pvx'] = {}
_mock_data['maxfwd'] = {}
_mock_data['registrar'] = {}
_mock_data['textops'] = {}
_mock_data['siputils'] = {}
_mock_data['xlog'] = {}
_mock_data['sanity'] = {}
_mock_data['kex'] = {}
_mock_data['corex'] = {}
_mock_data['jsonrpcs'] = {}
_mock_data['http_async_client'] = {}
_mock_data['xhttp'] = {}
_mock_data['websocket'] = {}
_mock_data['nathelper'] = {}
_mock_data['app_python3'] = {}
_mock_data['auth'] = {}
_mock_data['auth_ephemeral'] = {}
_mock_data['dispatcher'] = {}
_mock_data['uac'] = {}
_mock_data['permissions'] = {}
