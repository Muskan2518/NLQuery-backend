# analyze.py
import sys
import base64
from io import BytesIO
from pymongo import MongoClient
import pandas as pd
from ydata_profiling import ProfileReport
from weasyprint import HTML

if len(sys.argv) != 4:
    print("Usage: python analyze.py <mongo_url> <db_name> <collection_name>")
    sys.exit(1)

mongo_url = sys.argv[1]
db_name = sys.argv[2]
collection_name = sys.argv[3]

# Connect to MongoDB
client = MongoClient(mongo_url)
db = client[db_name]
collection = db[collection_name]

# Fetch data
data = list(collection.find())
df = pd.DataFrame(data)

if df.empty:
    print(f"Collection '{collection_name}' is empty.")
    sys.exit(0)

# Generate profiling report (HTML) in memory
report = ProfileReport(df, title=f"Auto Report - {collection_name}", explorative=True)
html_buffer = BytesIO()
report.to_file(html_buffer, output_format="html")
html_content = html_buffer.getvalue().decode("utf-8")

# Convert HTML to PDF in memory
pdf_buffer = BytesIO()
HTML(string=html_content).write_pdf(pdf_buffer)
pdf_bytes = pdf_buffer.getvalue()

# Encode PDF to base64
encoded_pdf = base64.b64encode(pdf_bytes).decode("utf-8")

# Output to stdout for Go to capture
print(encoded_pdf)
