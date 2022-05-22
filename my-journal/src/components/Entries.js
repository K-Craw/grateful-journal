import React from 'react';

const Notes = () => {
  return (
    <div>
        <form action="">
            <label for="user_name">Name:</label>
            <input id="user_name" name="user_name"></input>
            <label for="Entry">Entry:</label>
            <input id="Entry" name="Entry"></input>
            <input type="submit" value="Submit"/>
        </form>
    </div>
  );
};

export default Notes;