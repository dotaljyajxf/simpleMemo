import React, { Component } from "react";
import {HashRouter as Router, Route, Link, Switch} from 'react-router-dom'
import "../style/body.css";

import Calendar from '../components/calendar.jsx';
import Docs from '../components/docs.jsx';
import Memo from '../components/memo.jsx';

const zh = {
  hello: "Hello",
  login: "登录",
  message: "请输入信息"
};

class Body extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
        <Router >
          <Switch className="body-content">
            <Route exact name = "calednar" path="/cal" component={Calendar}></Route>
            <Route name = "docs" path="/docs" component={Docs}></Route>
            <Route name = "memo" path="/memo" component={Memo}></Route>
          </Switch>
        </Router>
     
    );
  }
}

export default Body;
