// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.73
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import org.threeten.bp.OffsetDateTime;

/**
 * V1Log
 */

public class V1Log {
  @SerializedName("timestamp")
  private OffsetDateTime timestamp = null;

  @SerializedName("node")
  private String node = null;

  @SerializedName("pod")
  private String pod = null;

  @SerializedName("container")
  private String container = null;

  @SerializedName("value")
  private String value = null;

  public V1Log timestamp(OffsetDateTime timestamp) {
    this.timestamp = timestamp;
    return this;
  }

   /**
   * Get timestamp
   * @return timestamp
  **/
  @ApiModelProperty(value = "")
  public OffsetDateTime getTimestamp() {
    return timestamp;
  }

  public void setTimestamp(OffsetDateTime timestamp) {
    this.timestamp = timestamp;
  }

  public V1Log node(String node) {
    this.node = node;
    return this;
  }

   /**
   * Get node
   * @return node
  **/
  @ApiModelProperty(value = "")
  public String getNode() {
    return node;
  }

  public void setNode(String node) {
    this.node = node;
  }

  public V1Log pod(String pod) {
    this.pod = pod;
    return this;
  }

   /**
   * Get pod
   * @return pod
  **/
  @ApiModelProperty(value = "")
  public String getPod() {
    return pod;
  }

  public void setPod(String pod) {
    this.pod = pod;
  }

  public V1Log container(String container) {
    this.container = container;
    return this;
  }

   /**
   * Get container
   * @return container
  **/
  @ApiModelProperty(value = "")
  public String getContainer() {
    return container;
  }

  public void setContainer(String container) {
    this.container = container;
  }

  public V1Log value(String value) {
    this.value = value;
    return this;
  }

   /**
   * Get value
   * @return value
  **/
  @ApiModelProperty(value = "")
  public String getValue() {
    return value;
  }

  public void setValue(String value) {
    this.value = value;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Log v1Log = (V1Log) o;
    return Objects.equals(this.timestamp, v1Log.timestamp) &&
        Objects.equals(this.node, v1Log.node) &&
        Objects.equals(this.pod, v1Log.pod) &&
        Objects.equals(this.container, v1Log.container) &&
        Objects.equals(this.value, v1Log.value);
  }

  @Override
  public int hashCode() {
    return Objects.hash(timestamp, node, pod, container, value);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Log {\n");
    
    sb.append("    timestamp: ").append(toIndentedString(timestamp)).append("\n");
    sb.append("    node: ").append(toIndentedString(node)).append("\n");
    sb.append("    pod: ").append(toIndentedString(pod)).append("\n");
    sb.append("    container: ").append(toIndentedString(container)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

