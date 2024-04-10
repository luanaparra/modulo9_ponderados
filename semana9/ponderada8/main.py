import time
import os
import pandas as pd

def process_file(file_path):
    start_time = time.time()

    df = pd.read_parquet(file_path)

    results = df.groupby('Station')['Temperature'].agg(['min', 'mean', 'max']).reset_index()

    results.sort_values(by='Station', inplace=True)

    print(results)

    end_time = time.time()
    execution_time = end_time - start_time
    print(f"Execution time: {execution_time:.2f} seconds")

def main():
    file_path = 'dados.parquet' 
    file_size = os.path.getsize(file_path)
    print(f"File size: {file_size / (1024 * 1024)} MB")

    process_file(file_path)

if __name__ == "__main__":
    main()
