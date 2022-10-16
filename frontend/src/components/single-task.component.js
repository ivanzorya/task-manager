import React, {useState} from "react";
import "bootstrap/dist/css/bootstrap.css";
import {Button, Card, Row, Col, InputGroup} from "react-bootstrap";

const Task = ({taskData, deleteSingleTask, changeSingleTask}) => {
    const [done, setDone] = useState(taskData.done)

    return (
        <Card>
            <Row>
                <Col>Subject:{ taskData !== undefined && taskData.subject}</Col>
                <Col><InputGroup.Checkbox checked={done} onChange={(event) => {
                    setDone(event.target.checked);
                    changeSingleTask({id: taskData._id, subject:taskData.subject, done: event.target.checked});
                    }}
                /></Col>
                <Col><Button onClick={() => deleteSingleTask(taskData._id)}>delete task</Button></Col>
            </Row>
        </Card>
    )
}

export default Task