import React, { useRef, useState } from "react";
function Post() {
  const baseURL = "http://localhost:8080";
  const id = useRef(0);
  const entry = useRef(null);
  const date = useRef(0);
  const user_name = useRef(null);
  const [postResult, setPostResult] = useState(null);
  const fortmatResponse = (res) => {
    return JSON.stringify(res, null, 2);
  }
  
  async function postData() {
    const postData = {
        id: id.current,
        entry: entry.current.value,
      date: date.current,
      user_name: user_name.current.value
    };
    try {
      const res = await fetch(`${baseURL}/entries`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "x-access-token": "token-value",
          "Access-Control-Allow-Origin": "http://localhost:8080/",
          "Access-Control-Allow-Headers": "Content-Type, Authorization"
        },
        body: JSON.stringify(postData),
      });
      if (!res.ok) {
        const message = `An error has occured: ${res.status} - ${res.statusText}`;
        throw new Error(message);
      }
      const data = await res.json();
      const result = {
        data: data,
      };
      setPostResult(fortmatResponse(result));
    } catch (err) {
      setPostResult(err.message);
    }
  }
  
  return (
    <div className="card">
      <div className="card-header">Post Here</div>
      <div className="card-body">
        <div className="form-group">
          <input type="text" className="form-control" ref={user_name} placeholder="Username" />
        </div>
        <div className="form-group">
          <input type="text" className="form-control" ref={entry} placeholder="Entry" />
        </div>
        <button className="btn btn-sm btn-primary" onClick={postData}>Submit</button>
        { postResult && <div className="alert alert-secondary mt-2" role="alert"><pre>{postResult}</pre></div> }
      </div>
    </div>
  );
}
export default Post;