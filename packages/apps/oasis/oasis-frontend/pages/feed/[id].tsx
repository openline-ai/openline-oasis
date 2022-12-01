import {useRouter} from "next/router";
import {Chat} from "./chat";

function CaseDetails() {
    const router = useRouter();
    const {id} = router.query;

    return (
        <Chat id={id}/>
    );
}

export default CaseDetails