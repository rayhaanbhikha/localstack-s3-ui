import React, { useEffect, useState } from 'react'
import { WrapTable, ResourceRow, BreadCrums } from '../../Components'
import { withRouter } from 'react-router-dom'

export const Component = ({history}) => {

    const [state, setState] = useState({
        path: "/",
        name: "localstack S3",
        type: "Root",
        breadcrums: [],
        resources: []
    })

    const fetchResources = async (resource) => {
        try {
            const res = await fetch(`http://localhost:8080/resource?path=${resource.path}`)
            const data = await res.json();
            setState(prevState => ({
                ...prevState,
                name: data.name,
                path: data.path,
                breadcrums: [
                    ...prevState.breadcrums,
                    { label: resource.name, url: resource.path }
                ],
                type: data.type,
                resources: data.children || []
            }))
        } catch (error) {
            console.log(error)
        }
    }

    useEffect(() => {
        fetchResources(state);
    }, [])

    const TableText = () => <>
        <BreadCrums history={history} breadcrums={state.breadcrums} />
        {state.type === "Root" ?
            <>
                <strong className="table-bucket-text">Buckets</strong>
            &nbsp;&nbsp;
                <strong className="table-bucket-nums">({state.resources.length})</strong>
            </>
            :
            <strong className="table-bucket-text">{state.name}</strong>
        }
    </>

    const TableBody = () =>
        state.resources.map((resource, index) =>
            <ResourceRow key={`bucketName-${index}`} resource={resource} fetchResources={fetchResources} />
        );

    return WrapTable(TableText, TableBody)
}
export const Resources = withRouter(Component)
