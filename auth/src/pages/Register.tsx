import React, { SyntheticEvent, useState } from 'react';
import {Redirect} from 'react-router-dom';
// Install package use to send http request
import axios from 'axios'
const Register = () => {
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [passwordConfirm, setPasswordConfirm] = useState('');
  const [redirect, setRedirect] = useState(false);

  const submit= async(e: SyntheticEvent) => {
    e.preventDefault();
    await axios.post('register', {
      "first_name": firstName,
      "last_name": lastName,
      "email": email,
      "password": password,
      "password_confirm": passwordConfirm
    });
    setRedirect(true);
  }
    if (redirect){
      return <Redirect to="/login"/>
    }
    return (
        <div>
        <form className="form-signin" onSubmit={submit}>
          <h1 className="h3 mb-3 font-weight-normal">Please sign up</h1>
          
          <input className="form-control" placeholder="First Name" onChange={e => setFirstName(e.target.value)} />
          
          <input className="form-control" placeholder="Last Name" onChange={e => setLastName(e.target.value)} />
          
          <input type="email" id="inputEmail" className="form-control" placeholder="Email address" required autoFocus onChange={e => setEmail(e.target.value)} />

          <input type="password" id="inputPassword" className="form-control" placeholder="Password" required onChange={e => setPassword(e.target.value)} />
          <input type="password" id="inputPassword" className="form-control" placeholder="Password Confirm" required onChange={e => setPasswordConfirm(e.target.value)} />
              <button className="btn btn-lg btn-primary btn-block" type="submit">Sign up</button>
        </form>
        </div>
    );
};

export default Register;