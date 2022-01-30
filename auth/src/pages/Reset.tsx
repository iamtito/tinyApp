import React, { SyntheticEvent, useState } from 'react';
import {Redirect} from 'react-router-dom';
import axios from 'axios';
const Reset = ({match}:{ match: any }) => {
    const [password, setPassword]=useState('');
    const [passwordConfirm, setPasswordConfirm]=useState('');
    const [redirect, setRedirect]=useState(false);
    
    const [notify, setNotify]=useState({
        show: false,
        error: false,
        message: ''
    });

    const submit = async (e: SyntheticEvent) =>{
        e.preventDefault();
        try {
            const token=match.params.token;

            await axios.post('reset', {
                token,
                password,
                password_confirm: passwordConfirm
            });
            
            setNotify({
                show: true,
                error: false,
                message: "password rest"
            });
            setRedirect(true);

            
        } catch (e) {
            setNotify({
                show: true,
                error: true,
                message: "Password reset failed"
            });
        }
    }
    let info;
    if (notify.show){
        info=(
            <div className={notify.error ? 'alert alert-danger' : 'alert alert-success'} role="alert">
                {notify.message}
            </div>
        )
    }
    if (redirect){
        return <Redirect to="/login" />
    }
    return (
        <div>
            <form className="form-signin" onSubmit={submit}>
                {info}
                <h1 className="h3 mb-3 font-weight-normal">Reset Password</h1>

                <input type="password" id="inputPassword" className="form-control" placeholder="Password" required onChange={e => setPassword(e.target.value)} />
                <input type="password" id="inputPassword" className="form-control" placeholder="Password Confirm" required onChange={e => setPasswordConfirm(e.target.value)} />

                <button className="btn btn-lg btn-primary btn-block" type="submit">Reset Password</button>
            </form>
        </div>
    );
};

export default Reset;