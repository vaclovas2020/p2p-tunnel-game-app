import logo from './../assets/images/appicon.png';

function Welcome() {
    return (
        <div className="bg-emerald-800 w-full h-full">
            <div className="flex flex-col justify-center items-center gap-4">
                <img className="max-w-[100%] w-40 h-auto" src={logo} alt="Logo" />
                <div className="bg-emerald-500 text-white text-2xl text-center">
                    Welcome to P2P Tunnel Game for Developers!
                </div>
            </div>
        </div>
    )
}

export default Welcome
