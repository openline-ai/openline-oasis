import {useRouter} from "next/router";
import {Menu} from "primereact/menu";


const LayoutMenu = () => {
    const router = useRouter();

    let items = [
        {
            label: 'Chats', icon: 'pi pi-mobile', command: () => {
                router.push('/feed');
            }
        }
    ];

    return (
        <Menu model={items}/>
    );
}

export default LayoutMenu
