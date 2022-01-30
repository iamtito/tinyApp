import React from 'react';

const Home = ({user}:{user: any}) => {
    // const [message, setMessage] = useState('')
    // setMessage()
    // setMessage()
    let message;
    if (user){
        message=`Hi ${user.first_name} ${user.last_name}`;
    }else{
        message=`you are not logged in`;
    }
    return (
        <div className="container">
            <h1>{message}</h1>
        </div>
    );
};

export default Home;