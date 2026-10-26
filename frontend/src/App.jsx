import {useEffect, useState} from "react";
import './App.css'

const API = "http://localhost:8080"

async function request(path, options) {
  const res = await fetch(`${API}/${path}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...options?.headers,
    },
  });
  if (!res.ok) throw new Error(`HTTP ${res.status}`);
  return res.status === 204 ? null : res.json();
}

function TypesView() {
  const [types, setTypes] = useState([]);

  const [typeId, setTypeId] = useState(0);
  const [typeName, setTypeName] = useState("");

  async function loadData() {
    const [t] = await Promise.all([
      request("types"),
    ]);

    setTypes(t);
  }

  useEffect(() => {
    loadData().catch(console.error);
  }, []);

  async function addType(e) {
    e.preventDefault();
    await request("types" , {
      method: "PUT",
      body: JSON.stringify({"id":typeId, "name":typeName}),
    });
    setTypeName("");
    setTypeId(typeId + 1);
  }

  return (
      <section>
        <h2>Types</h2>

        <form onSubmit={addType}>
          <input
              type="text"
              placeholder="Type Name"
              value={typeName}
              onChange={(e) => setTypeName(e.target.value)}
              required
          />
          <input
              type="number"
              placeholder="Type ID"
              value={typeId}
              onChange={(e) => setTypeId(parseInt(e.target.value, 10))}
              required
          />
          <button>Add type</button>
        </form>
        <table>
          <thead>
          <tr><td>Id</td><td>Name</td></tr>
          </thead>
          <tbody>
          {types.map((type) => (
              <tr key={type.id}>
                <td>{type.id}</td>
                <td>{type.name}</td>
              </tr>
          ))}
          </tbody>
        </table>
      </section>
  )
}
function SubTypesView() {
  const [subTypes, setSubTypes] = useState([]);


  const [subTypeId, setSubTypeId] = useState(0);
  const [subTypeName, setSubTypeName] = useState("");

  async function loadData() {
    const [st ] = await Promise.all([
      request("subtypes"),
    ]);

    setSubTypes(st);
  }
  useEffect(() => {
    loadData().catch(console.error);
  }, []);

  async function addSubType(e) {
    e.preventDefault();
    await request("subTypes", {
      method: "PUT",
      body: JSON.stringify({"id":subTypeId, "name":subTypeName}),
    });
    setSubTypeName("");
    setSubTypeId(subTypeId +1)
  }

  return (
      <section>
        <h2> Subtypes</h2>
        <form onSubmit={addSubType}>
          <input
            type="text"
            placeholder="Subtype Name"
            defaultValue={subTypeName}
            onChange={(e) => setSubTypeName(e.target.value)}
            required
          />
          <input
            type="number"
            placeholder="Subtype ID"
            defaultValue={subTypeId}
            onChange={(e) => setSubTypeId(parseInt(e.target.value,10))}
            required
            />
          <button>Add subtype</button>
        </form>
        <table>
          <thead>
          <tr><td>Id</td><td>Name</td></tr>
          </thead>
          <tbody>
          {subTypes.map((subType) => (
              <tr key={`Sub-${subType.id}`}>
                <td>{subType.id}</td>
                <td>{subType.name}</td>
              </tr>
          ))}
          </tbody>
        </table>
      </section>
  )
}
function App() {
  const [view, setView] = useState("types");
 return (
  <>
    <header className="App-header">
      <div className="header-inner">
        <h1> Admin Dashboard</h1>

        <nav>
          <button
              className={view === "types" ? "nav-btn active": "nav-btn"}
              onClick={() => setView("types")}
            >
            Types
          </button>
          <button
            className={view === "subTypes" ? "nav-btn active": "nav-btn"}
            onClick={() => setView("subTypes")}
            >
            SubTypes
          </button>
        </nav>
      </div>
    </header>
    <main>
      {view === "types" ? <TypesView/> : <SubTypesView/>}
    </main>
  </>
 )
}

export default App;

