import { useState, useEffect } from 'react';
import axios from 'axios';
export const BASE_URL=import.meta.env.MODE==="development"?"http://127.0.0.1:4000/api":"/api";
function App() {

  const backgroundAnimation = {
    background: "linear-gradient(-45deg, #FF7E5F, #FEB47B, #6A11CB, #2575FC)",
    backgroundSize: "400% 400%",
    animation: "gradient 10s ease infinite",
    height: "100vh",
    width: "100vw",
    padding: "0",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  };

  // Animation keyframes using a style tag
  const style = `
    @keyframes gradient {
      0% { background-position: 0% 50%; }
      50% { background-position: 100% 50%; }
      100% { background-position: 0% 50%; }
    }
  `;
  const tableStyle = {
    width: '100%',
    borderCollapse: 'collapse',
    fontFamily: 'Arial, sans-serif',
    backgroundColor: '#f8f9fa',
  };

  const thStyle = {
    backgroundColor: '#4A90E2',
    color: '#fff',
    textAlign: 'left',
    padding: '12px',
    borderBottom: '2px solid #e0e0e0',
    position: 'sticky',
    top: 0,
    zIndex: 1,
  };

  const tdStyle = {
    padding: '8px',
    textAlign: 'left',
    borderBottom: '1px solid #e0e0e0',
    fontSize: '14px',
  };

  const rowHoverStyle = {
    transition: 'background-color 0.3s ease',
    cursor: 'pointer',
  };

  const buttonStyle = {
    padding: '6px 12px',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
    marginRight: '10px',
    transition: 'background-color 0.3s ease',
  };

  const updateButtonStyle = {
    ...buttonStyle,
    background: "linear-gradient(135deg, #FF7E5F 0%, #feb47b 100%)",
    color: '#fff',
  };

  const deleteButtonStyle = {
    ...buttonStyle,
    background: "linear-gradient(135deg, #6a11cb 0%, #2575fc 100%)",
    color: '#fff',
  };

  const inputStyle = {
    width: '60%',
    maxWidth: '400px',
    padding: '12px 16px',
    margin: '20px 10px',
    border: '2px solid #ddd',
    borderRadius: '25px',
    outline: 'none',
    fontSize: '16px',
    color: '#333',
    backgroundColor: '#f9f9f9',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
  };

  const buttonStyle1 = {
    backgroundColor: '#4A90E2',
    color: '#fff',
    padding: '10px 16px',
    border: 'none',
    borderRadius: '20px',
    cursor: 'pointer',
    fontSize: '16px',
    fontWeight: 'bold',
    transition: 'background-color 0.3s ease, transform 0.2s ease',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
    margin: '10px',
  };

  const [todo, setTodo] = useState({ body: "", completed: false });
  const [todos, setTodos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchTodos = async () => {
    try {
      const response = await axios.get(`${BASE_URL}/gettodos`);
      setTodos(response.data);
    } catch (err) {
      setError('Failed to fetch TODOs');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTodos();
  }, [loading]);

  const handleChange = (e) => {
    setTodo({ ...todo, body: e.target.value });
  };

  const submitTodo = async () => {
    if (todo.body === "") {
      alert("Please enter a TODO");
    } else {
      try {
        setLoading(true);
        const res = await axios.post(`${BASE_URL}/createtodo`, todo);
        setTodo({ ...todo, body: "" });
        setLoading(false);
        fetchTodos();
      } catch (error) {
        console.error("Error submitting todo:", error);
      }
    }
  };

  const deleteTodo = async (id) => {
    try {
      setLoading(true);
      await axios.delete(`${BASE_URL}/deletetodo/${id}`);
      setLoading(false);
      fetchTodos();
    } catch (error) {
      console.error("Error deleting todo:", error);
    }
  };

  const updateTodo = async (id) => {
    try {
      setLoading(true);
      await axios.patch(`${BASE_URL}/updatetodo/${id}`);
      setLoading(false);
      fetchTodos();
    } catch (error) {
      console.error("Error updating todo:", error);
    }
  };

  return (
    <>
    <style>{style}</style>
 <div style={backgroundAnimation}>
      <div style={{
        width: '95%',
        maxWidth: '800px',
        background: "linear-gradient(135deg, #FF7E5F 0%, #feb47b 100%)",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        padding: '20px',
        borderRadius: "10px",
        marginBottom: '20px',
      }}>
        <h2 style={{ color: 'purple' }}>My Todo App</h2>
        <input
          type="text"
          placeholder="Enter your Todo..."
          style={inputStyle}
          value={todo.body}
          onChange={handleChange}
        />
        <button style={buttonStyle1} onClick={submitTodo}>Enter Todo</button>
      </div>

      <div style={{
        width: '95%',
        maxWidth: '800px',
        background: 'white',
        borderRadius: '10px',
        overflowY: 'scroll',
        maxHeight: '60vh',
      }}>
        <table style={tableStyle}>
          <thead>
            <tr>
              <th style={thStyle}>No</th>
              <th style={thStyle}>TODO</th>
              <th style={thStyle}>Status</th>
              <th style={thStyle}>Update</th>
              <th style={thStyle}>Delete</th>
            </tr>
          </thead>
          <tbody>
            {todos&&todos.map((todo, index) => (
              <tr key={index} style={rowHoverStyle}>
                <td style={tdStyle}>{index + 1}</td>
                <td style={tdStyle}>{todo.body}</td>
                <td style={tdStyle}>{todo.completed ? "Completed" : "Pending"}</td>
                <td style={tdStyle}>
                {todo.completed?<p style={{color:"green"}}>Done!</p>:<button onClick={() => updateTodo(todo.id)} style={updateButtonStyle}>Update</button>}  
                </td>
                <td style={tdStyle}>
                  <button onClick={() => deleteTodo(todo.id)} style={deleteButtonStyle}>Delete</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
    </>
   
  );
}

export default App;
