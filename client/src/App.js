import './App.css';

import * as React from "react";
import { useForm } from "react-hook-form";
import { Link } from "react-router-dom";
import axios from 'axios';
const baseURL = "https://api.local:446/users";
export default function App() {
  axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';
  const { register, handleSubmit } = useForm();
  const onSubmit = (data) => alert(JSON.stringify(data));


  const [post, setPost] = React.useState(null);

  React.useEffect(() => {
    axios.get(baseURL).then((response) => {
      setPost(response.data);
    });
  }, []);

  if (!post) return null;
  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <h1>{post.title}</h1>
      <input {...register("firstName", { required: true })} placeholder="First name" />

      <input {...register("lastName", { minLength: 2 })} placeholder="Last name" />

      <select {...register("category")}>
        <option value="">Select...</option>
        <option value="A">Category A</option>
        <option value="B">Category B</option>
      </select>
      <Link to="/invoices">Invoices</Link> |{" "}
      <Link to="/expenses">Expenses</Link>
      <input type="submit" />
    </form>
  );
}

