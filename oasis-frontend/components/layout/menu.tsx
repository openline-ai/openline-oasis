import {useRouter} from "next/router";
import {Menu} from "primereact/menu";


const LayoutMenu = () => {
    const router = useRouter();

    let items = [
        {
            label: 'Cases', icon: 'pi pi-mobile', command: () => {
                router.push('/case');
            }
        }
    ];

    return (
        <Menu model={items}/>
    );
}

export default LayoutMenu
