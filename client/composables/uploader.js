export const FormImgUploader = async (formData) => {
    if (!formData.get('photo').name) {
        return { data: null }
    }
    let body = new FormData()
    body.append('file', formData.get('photo'))
    return await fetch("/api/upload", {
        method: "POST",
        body: body,
    }).then(response => response.json()).catch(err => ({ errors: err }))

}