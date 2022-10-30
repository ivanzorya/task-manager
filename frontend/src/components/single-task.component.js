import React, {useState} from "react";
import "bootstrap/dist/css/bootstrap.css";
import {Button, Card, Row, Col, InputGroup} from "react-bootstrap";

const Task = ({taskData, deleteSingleTask, changeSingleTask}) => {
    const [done, setDone] = useState(taskData.done)

    return (
        <Card className="stndrt-class">
            <Row>
                <Col sm={9}>Subject: { taskData !== undefined && taskData.subject}</Col>
                <Col sm={1}>
                    <InputGroup.Checkbox checked={done} onChange={(event) => {
                            setDone(event.target.checked);
                            changeSingleTask({id: taskData._id, subject:taskData.subject, done: event.target.checked});
                        }}
                    />
                </Col>
                <Col sm={2}>
                    <Button className="stndrt-class" onClick={() => deleteSingleTask(taskData._id)}>delete task</Button>
                </Col>
            </Row>
        </Card>
    )
}

export default Task