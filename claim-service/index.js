import Express from "express";

const app = Express();
const port = 4000;

const claim = [
    { docName: 1, name: 'John Doe' },
    { docName: 2, name: 'Jane Smith' },
    { docName: 3, name: 'Bob Johnson' },
];

app.get('/api/v1', (req, res) => {
    res.json(claim);
})

app.get('/api/v1/:id', (req, res) => {
    const matchClaim = claim.find(u => u.docName === parseInt(req.params.id));
    if (!matchClaim) return res.status(404).send('Data not found');
    res.json(matchClaim);
});

app.listen(port, ()=>{
    console.log(`Listening on port ${port}`);
})