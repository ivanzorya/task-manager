import React, {useState, useEffect} from "react";
import axios from "axios";
import {Button, Form, Container, Modal, InputGroup } from "react-bootstrap"
import Task from "./single-task.component";

const Tasks = () => {

    const [tasks, setTasks] = useState([])
    const [addNewTask, setAddNewTask] = useState(false)
    const [newTask, setNewTask] = useState({"subject": "", "done": false})

    const changeSingleTask = (updatedTask) => {
        console.log(updatedTask)
        var url = "http://localhost:5000/task/update/" + updatedTask.id
        axios.put(url, updatedTask)
            .then(response => {
            if(response.status === 200){
                getAllTasks()
            }
        })
    }

    const addSingleTask = () => {
        setAddNewTask(false)
        var url = "http://localhost:5000/task/create"
        axios.post(url, {
            "subject": newTask.subject,
            "done": newTask.done,
        }).then(response => {
            if(response.status === 200){
                getAllTasks()
            }
        })
    }

    const getAllTasks = () => {
        var url = "http://localhost:5000/tasks"
        axios.get(url, {
            responseType: "json"
        }).then(response => {
            if(response.status === 200){
                setTasks(response.data);
            }
        })
    }

    const deleteSingleTask = (id) => {
        var url = "http://localhost:5000/task/delete/" + id
        axios.delete(url, {
        }).then(response => {
            if(response.status === 200){
                getAllTasks();
            }
        })
    }

    useEffect(() => {
        getAllTasks();
    }, [])
 
    return (
        <div>
            
            <Container>
                <Button onClick={() => setAddNewTask(true)}>Add new task</Button>
            </Container>

            <Container>
                {tasks != null && tasks.map((task) => (
                    <Task key={task._id} taskData={task} deleteSingleTask={deleteSingleTask} changeSingleTask={changeSingleTask}/>
                ))}
            </Container>
            
            <Modal show={addNewTask} onHide={() => setAddNewTask(false)} centered>
                <Modal.Header closeButton>
                    <Modal.Title>Add Task</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <Form.Group>
                        <Form.Control onChange={(event) => {newTask.subject = event.target.value}}/>
                        <InputGroup.Checkbox type="checkbox" onChange={(event) => {newTask.done = event.target.checked}}/>
                    </Form.Group>
                    <Button onClick={() => addSingleTask()}>Add</Button>
                    <Button onClick={() => setAddNewTask(false)}>Cancel</Button>
                </Modal.Body>
            </Modal>
        </div>
    );
}

export default Tasks